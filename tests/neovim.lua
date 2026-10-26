local hea = require("hea")
local project = vim.env.HEAD_TEST_PROJECT or vim.fn.tempname()
local state = vim.env.HEAD_TEST_DATA or vim.fn.tempname()
vim.fn.mkdir(project, "p")
vim.fn.mkdir(state, "p")
local bin = assert(vim.env.HEAD_TEST_BIN, "HEAD_TEST_BIN must point to the built hea executable")
hea.setup({ cmd = bin, project = project, data_dir = state })
vim.o.lines, vim.o.columns = 80, 160
local notices = {}
vim.notify = function(message) table.insert(notices, message) end

local function wait(predicate, message)
  assert(vim.wait(5000, predicate, 10), message or "timed out")
end
local function request(args, expect_error)
  local done, value, failure = false, nil, nil
  hea.request(args, function(result, err) value, failure, done = result, err, true end)
  wait(function() return done end, "request timed out")
  assert(expect_error and failure or not expect_error and not failure, failure)
  return value
end
local function contents(buf)
  return table.concat(vim.api.nvim_buf_get_lines(buf, 0, -1, false), "\n")
end
local function key(view, mapping)
  vim.api.nvim_set_current_win(view.win)
  local lhs = vim.api.nvim_replace_termcodes(mapping, true, false, true)
  vim.api.nvim_feedkeys(lhs, "xt", false)
end
local function close(view)
  if vim.api.nvim_buf_is_valid(view.buf) then vim.api.nvim_buf_delete(view.buf, { force = true }) end
end

if vim.env.HEAD_TEST_REOPEN == "1" then
  assert(#request({ "ticket", "list" }) == 2, "tickets lost across editor restart")
  assert(#request({ "notification", "list" }) == 4, "notifications lost across editor restart")
  print("HEAD Neovim restart persistence: passed")
  return
end

local first = request({ "ticket", "create", "--title", "First UI ticket", "--criterion", "Reviewed" })
local board = hea.board()
wait(function() return contents(board.buf):find("First UI ticket", 1, true) end, "board did not load")
assert(contents(board.buf):find("IN_PROGRESS", 1, true))
vim.api.nvim_win_set_cursor(board.win, { 4, 0 })
vim.ui.select = function(_, _, callback) callback("in_progress") end
key(board, "m")
wait(function() return #(board.items.in_progress or {}) == 1 end, "move mapping did not refresh")
assert(request({ "ticket", "get", tostring(first.id) }).state == "in_progress")

vim.ui.input = function(_, callback) callback("Created with a") end
key(board, "a")
wait(function() return #(board.items.todo or {}) == 1 end, "create mapping failed")
vim.api.nvim_win_set_cursor(board.win, { 4, 0 })
vim.ui.select = function(_, _, callback) callback("Description") end
key(board, "e")
local editor = vim.api.nvim_get_current_buf()
vim.api.nvim_buf_set_lines(editor, 0, -1, false, { "Multiline description", "Café 🚀" })
vim.api.nvim_feedkeys(vim.api.nvim_replace_termcodes("<C-s>", true, false, true), "xt", false)
wait(function() return not vim.api.nvim_buf_is_valid(editor) end, "editor did not save")
assert(request({ "ticket", "get", tostring(board.items.todo[1].id) }).description == "Multiline description\nCafé 🚀")

request({ "notification", "add", "--text", "Manual review", "--priority", "3" })
local queue = hea.notifications()
wait(function() return #queue.items == 1 end, "queue did not load")
vim.api.nvim_win_set_cursor(queue.win, { 3, 0 })
key(queue, "u")
wait(function() return queue.items[1].priority == 2 end, "requeue failed")
key(queue, "x")
wait(function() return #queue.items == 0 end, "dismiss failed")

local output = hea.run("OK")
wait(function() return not hea.running() end, "agent did not finish")
assert(contents(output.buf):find("fixture response", 1, true), contents(output.buf))
assert(contents(output.buf):find("Task completed", 1, true))
close(output)
local items = request({ "notification", "list" })
assert(#items == 1 and items[1].text == "Codex task completed")

output = hea.run("FAIL")
wait(function() return not hea.running() end, "failed agent did not finish")
assert(contents(output.buf):find("fixture failure", 1, true))
close(output)
items = request({ "notification", "list" })
assert(#items == 2 and items[1].text:find("failed", 1, true))

output = hea.run("SLOW")
wait(function() return contents(output.buf):find("Codex is working", 1, true) end, "output was not streamed")
assert(hea.running())
hea.stop()
wait(function() return not hea.running() end, "stop did not cancel")
assert(contents(output.buf):find("Task cancelled", 1, true))
close(output)
items = request({ "notification", "list" })
assert(#items == 3)

-- Editor shutdown uses the same real process cleanup path.
output = hea.run("SLOW shutdown")
wait(function() return contents(output.buf):find("Codex is working", 1, true) end)
hea.shutdown()
wait(function() return not hea.running() end)
close(output)

-- Invalid commands and missing executables are surfaced without breaking views.
request({ "ticket", "get", "999999" }, true)
hea.setup({ cmd = "/does-not-exist" })
request({ "ticket", "list" }, true)
hea.setup({ cmd = bin })
close(board)
board = hea.board()
wait(function() return #(board.items.in_progress or {}) == 1 end)
close(board)
close(queue)
assert(#notices > 0)
if not vim.env.HEAD_TEST_PROJECT then vim.fn.delete(project, "rf") end
if not vim.env.HEAD_TEST_DATA then vim.fn.delete(state, "rf") end
print("HEAD Neovim integration: passed")
