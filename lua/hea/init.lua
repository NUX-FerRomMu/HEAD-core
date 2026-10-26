local M = {}
local config = { cmd = "hea" }
local active
local views = {}
local states = { "todo", "in_progress", "review", "done" }

local function notify(message, level)
  vim.notify("HEAD: " .. message, level or vim.log.levels.INFO)
end

local function argv(args, project)
  local command = { config.cmd, "--project", project or config.project or vim.fn.getcwd() }
  if config.data_dir then vim.list_extend(command, { "--data-dir", config.data_dir }) end
  return vim.list_extend(command, args)
end

-- All durable state and business decisions stay in Go.
function M.request(args, callback, project)
  local ok, job = pcall(vim.system, argv(args, project), { text = true }, function(result)
    vim.schedule(function()
      if result.code ~= 0 then
        local err = vim.trim(result.stderr or "")
        if err == "" then err = "command exited with code " .. result.code end
        notify(err, vim.log.levels.ERROR)
        if callback then callback(nil, err) end
        return
      end
      local decoded, value = pcall(vim.json.decode, result.stdout)
      if not decoded then
        notify("invalid response from hea", vim.log.levels.ERROR)
        if callback then callback(nil, "invalid JSON") end
        return
      end
      if callback then callback(value) end
    end)
  end)
  if not ok then
    notify(tostring(job), vim.log.levels.ERROR)
    if callback then callback(nil, tostring(job)) end
  end
  return ok and job or nil
end

local function lines(text)
  return vim.split(text or "", "\n", { plain = true })
end

local function write(buf, content)
  if not vim.api.nvim_buf_is_valid(buf) then return end
  vim.bo[buf].modifiable = true
  vim.api.nvim_buf_set_lines(buf, 0, -1, false, content)
  vim.bo[buf].modifiable = false
end

local function panel(name, project)
  vim.cmd("botright new")
  local buf, win = vim.api.nvim_get_current_buf(), vim.api.nvim_get_current_win()
  vim.bo[buf].buftype = "nofile"
  vim.bo[buf].bufhidden = "wipe"
  vim.bo[buf].swapfile = false
  vim.bo[buf].filetype = "hea"
  vim.wo[win].wrap = false
  local view = { buf = buf, win = win, project = project or config.project or vim.fn.getcwd(), items = {}, generation = 0, kind = name }
  views[buf] = view
  vim.api.nvim_create_autocmd("BufWipeout", { buffer = buf, once = true, callback = function() views[buf] = nil end })
  vim.api.nvim_buf_set_name(buf, "HEAD://" .. name .. "/" .. buf)
  vim.keymap.set("n", "q", function() vim.api.nvim_buf_delete(buf, { force = true }) end, { buffer = buf, silent = true })
  return view
end

local function map(view, key, action, modes)
  vim.keymap.set(modes or "n", key, action, { buffer = view.buf, silent = true })
end

local function refresh(view, args, render)
  view.generation = view.generation + 1
  local generation = view.generation
  M.request(args, function(value, err)
    if not err and generation == view.generation and vim.api.nvim_buf_is_valid(view.buf) and vim.api.nvim_win_is_valid(view.win) then render(value) end
  end, view.project)
end

local function mutate(view, args)
  M.request(args, function(_, err) if not err then view.refresh() end end, view.project)
end

local function fit(text, width)
  text = text:gsub("[\r\n\t]", " ")
  while vim.fn.strdisplaywidth(text) > width do text = vim.fn.strcharpart(text, 0, vim.fn.strchars(text) - 1) end
  return text .. string.rep(" ", math.max(0, width - vim.fn.strdisplaywidth(text)))
end

local function ticket_details(ticket, project)
  local view = panel("ticket-" .. ticket.id, project)
  local content = { "#" .. ticket.id .. " " .. ticket.title, "State: " .. ticket.state, "" }
  vim.list_extend(content, lines(ticket.description))
  vim.list_extend(content, { "", "Acceptance criteria:" })
  for _, criterion in ipairs(ticket.criteria) do vim.list_extend(content, lines("- " .. criterion)) end
  write(view.buf, content)
end

local function edit_field(view, ticket)
  vim.ui.select({ "Title", "Description", "Acceptance criteria" }, { prompt = "Edit ticket #" .. ticket.id }, function(field)
    if not field then return end
    if field == "Title" then
      vim.ui.input({ prompt = "Title: ", default = ticket.title }, function(value)
        if value then mutate(view, { "ticket", "edit", tostring(ticket.id), "--title", value }) end
      end)
      return
    end
    local editor = panel("edit-ticket-" .. ticket.id, view.project)
    local content = field == "Description" and lines(ticket.description) or ticket.criteria
    write(editor.buf, content)
    vim.bo[editor.buf].modifiable = true
    vim.bo[editor.buf].filetype = "markdown"
    local saving = false
    map(editor, "<C-s>", function()
      if saving then return end
      local values = vim.api.nvim_buf_get_lines(editor.buf, 0, -1, false)
      local args = { "ticket", "edit", tostring(ticket.id) }
      if field == "Description" then
        vim.list_extend(args, { "--description", table.concat(values, "\n") })
      else
        for _, value in ipairs(values) do if value ~= "" then vim.list_extend(args, { "--criterion", value }) end end
        if #args == 3 then table.insert(args, "--clear-criteria") end
      end
      saving = true
      M.request(args, function(_, err)
        saving = false
        if not err then
          if vim.api.nvim_buf_is_valid(editor.buf) then vim.api.nvim_buf_delete(editor.buf, { force = true }) end
          view.refresh()
        end
      end, view.project)
    end, { "n", "i" })
    notify("Edit " .. field:lower() .. "; Ctrl-S saves, q cancels. Criteria use one line each.")
  end)
end

function M.board()
  local view = panel("board")
  view.refresh = function()
    refresh(view, { "ticket", "board" }, function(board)
      local width = math.max(16, math.floor((vim.api.nvim_win_get_width(view.win) - 9) / 4))
      view.width, view.items = width, board
      local headings, count = {}, 0
      for _, state in ipairs(states) do
        table.insert(headings, fit(state:upper(), width))
        count = math.max(count, #board[state])
      end
      local content = { "a create | e edit | m move | Enter details | x run | r refresh | q close", table.concat(headings, " | "), string.rep("-", width * 4 + 9) }
      for row = 1, math.max(1, count) do
        local cells = {}
        for _, state in ipairs(states) do
          local ticket = board[state][row]
          table.insert(cells, fit(ticket and ("#" .. ticket.id .. " " .. ticket.title) or "", width))
        end
        table.insert(content, table.concat(cells, " | "))
      end
      write(view.buf, content)
    end)
  end
  local function selected()
    local row = vim.api.nvim_win_get_cursor(view.win)[1] - 3
    local column = math.min(4, math.floor((vim.fn.virtcol(".") - 1) / ((view.width or 16) + 3)) + 1)
    return (view.items[states[column]] or {})[row]
  end
  map(view, "r", view.refresh)
  map(view, "a", function()
    vim.ui.input({ prompt = "Ticket title: " }, function(value)
      if value then mutate(view, { "ticket", "create", "--title", value }) end
    end)
  end)
  map(view, "<CR>", function() local ticket = selected(); if ticket then ticket_details(ticket, view.project) end end)
  map(view, "e", function() local ticket = selected(); if ticket then edit_field(view, ticket) end end)
  map(view, "m", function()
    local ticket = selected()
    if ticket then vim.ui.select(states, { prompt = "Move #" .. ticket.id .. " to:" }, function(state)
      if state then mutate(view, { "ticket", "move", tostring(ticket.id), state }) end
    end) end
  end)
  map(view, "x", function()
    local ticket = selected()
    if ticket then
      M.run("Work on ticket #" .. ticket.id .. ": " .. ticket.title .. "\n\n" .. ticket.description .. "\n\nAcceptance criteria:\n- " .. table.concat(ticket.criteria, "\n- "), view.project)
    end
  end)
  view.refresh()
  return view
end

function M.notifications()
  local view = panel("notifications")
  view.refresh = function()
    refresh(view, { "notification", "list" }, function(items)
      view.items = items
      local content = { "a add | n next | Enter preview | x dismiss | u requeue | r refresh | q close", "" }
      for _, item in ipairs(items) do
        table.insert(content, ("#%d [P%d] %s (%s)"):format(item.id, item.priority, item.text:gsub("[\r\n]", " "), item.source))
      end
      if #items == 0 then table.insert(content, "No pending notifications.") end
      write(view.buf, content)
    end)
  end
  local function selected() return view.items[vim.api.nvim_win_get_cursor(view.win)[1] - 2] end
  map(view, "r", view.refresh)
  map(view, "a", function() vim.ui.input({ prompt = "Notification: " }, function(value)
    if value then mutate(view, { "notification", "add", "--text", value }) end
  end) end)
  map(view, "n", function()
    M.request({ "notification", "next" }, function(item, err)
      if err or item == vim.NIL or not vim.api.nvim_win_is_valid(view.win) then return end
      for row, existing in ipairs(view.items) do if existing.id == item.id then vim.api.nvim_win_set_cursor(view.win, { row + 2, 0 }); break end end
    end, view.project)
  end)
  map(view, "<CR>", function()
    local item = selected()
    if item then local preview = panel("notification-" .. item.id, view.project); write(preview.buf, lines(item.text)) end
  end)
  for key, operation in pairs({ x = "dismiss", u = "requeue" }) do
    map(view, key, function() local item = selected(); if item then mutate(view, { "notification", operation, tostring(item.id) }) end end)
  end
  view.refresh()
  return view
end

local function append(buf, text)
  if not vim.api.nvim_buf_is_valid(buf) then return end
  vim.bo[buf].modifiable = true
  vim.api.nvim_buf_set_lines(buf, -1, -1, false, lines(text))
  vim.bo[buf].modifiable = false
end

local function event_text(line)
  local ok, event = pcall(vim.json.decode, line)
  if not ok or type(event) ~= "table" then return line end
  local labels = {
    ["hea.run.started"] = "Task started", ["hea.run.completed"] = "Task completed",
    ["hea.run.failed"] = "Task failed", ["hea.run.cancelled"] = "Task cancelled",
    ["turn.started"] = "Codex is working", ["turn.completed"] = "Codex finished",
  }
  if labels[event.type] then return labels[event.type] .. (event.error and (": " .. tostring(event.error)) or "") end
  local item = event.item
  if item then return item.text or item.aggregated_output or item.command or ("[" .. (item.type or "item") .. "]") end
  if event.error then return type(event.error) == "table" and event.error.message or tostring(event.error) end
  return event.message or ("[" .. (event.type or "event") .. "]")
end

function M.run(prompt, project)
  if active then notify("A task is already running; use :HEADStop first", vim.log.levels.WARN); return end
  if not prompt or prompt == "" then
    vim.ui.input({ prompt = "Codex task: " }, function(value) if value and value ~= "" then M.run(value, project) end end)
    return
  end
  local view = panel("agent", project)
  write(view.buf, { "Codex task — :HEADStop cancels", prompt, "" })
  local args = { "agent", "run", "--prompt", prompt }
  if config.model then vim.list_extend(args, { "--model", config.model }) end
  local pending = ""
  local run = { view = view }
  active = run
  local ok, job = pcall(vim.system, argv(args, view.project), {
    text = true,
    stdout = function(err, data)
      if err then vim.schedule(function() append(view.buf, tostring(err)) end) end
      if not data then return end
      vim.schedule(function()
        pending = pending .. data
        while true do
          local newline = pending:find("\n", 1, true)
          if not newline then break end
          local line = pending:sub(1, newline - 1)
          pending = pending:sub(newline + 1)
          if line ~= "" then append(view.buf, event_text(line)) end
        end
      end)
    end,
    stderr = function(_, data) if data then vim.schedule(function() append(view.buf, vim.trim(data)) end) end end,
  }, function(result)
    vim.schedule(function()
      if pending ~= "" then append(view.buf, event_text(pending)) end
      append(view.buf, "[hea exited " .. result.code .. "]")
      if active == run then active = nil end
      for _, open_view in pairs(views) do if open_view.kind == "notifications" and open_view.refresh then open_view.refresh() end end
      notify(result.code == 0 and "Task completed; see :HEADNotifications" or "Task stopped or failed; see output and :HEADNotifications", result.code == 0 and vim.log.levels.INFO or vim.log.levels.WARN)
    end)
  end)
  if not ok then active = nil; append(view.buf, tostring(job)); notify(tostring(job), vim.log.levels.ERROR); return end
  run.job = job
  return view
end

function M.running() return active ~= nil end
function M.stop() if active and active.job then active.job:kill(15) end end
function M.shutdown()
  if active and active.job then
    local job = active.job
    job:kill(15)
    job:wait(5000)
  end
end

function M.status()
  M.request({ "agent", "status" }, function(value, err)
    if not err then notify((value.version or "Codex unavailable") .. ": " .. (value.detail or "")) end
  end)
end

function M.setup(options)
  config = vim.tbl_extend("force", config, options or {})
  local group = vim.api.nvim_create_augroup("HEAD", { clear = true })
  vim.api.nvim_create_autocmd("VimLeavePre", { group = group, callback = M.shutdown })
end

M.setup()
return M
