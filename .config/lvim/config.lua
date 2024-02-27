-- Read the docs: https://www.lunarvim.org/docs/configuration
-- Video Tutorials: https://www.youtube.com/watch?v=sFA9kX-Ud_c&list=PLhoH5vyxr6QqGu0i7tt_XoVK9v-KvZ3m6
-- Forum: https://www.reddit.com/r/lunarvim/
-- Discord: https://discord.com/invite/Xb9B4Ny

lvim.transparent_window = true

lvim.plugins = {
  {"folke/tokyonight.nvim"},
  {"folke/trouble.nvim"}
}

lvim.colorscheme = "tokyonight-storm"

-- move between buffers
lvim.keys.normal_mode["<S-h>"] = ":BufferLineCyclePrev<CR>"
lvim.keys.normal_mode["<S-l>"] = ":BufferLineCycleNext<CR>"

lvim.keys.normal_mode["tr"] = ":TroubleToggle<CR>"

vim.opt.relativenumber = true
