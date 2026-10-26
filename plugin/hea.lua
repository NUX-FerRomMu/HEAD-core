if vim.g.loaded_hea then return end
vim.g.loaded_hea = true

local function hea() return require("hea") end
vim.api.nvim_create_user_command("HEADBoard", function() hea().board() end, {})
vim.api.nvim_create_user_command("HEADNotifications", function() hea().notifications() end, {})
vim.api.nvim_create_user_command("HEADRun", function(opts) hea().run(opts.args) end, { nargs = "*" })
vim.api.nvim_create_user_command("HEADStop", function() hea().stop() end, {})
vim.api.nvim_create_user_command("HEADStatus", function() hea().status() end, {})
