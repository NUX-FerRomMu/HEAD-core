#!/bin/sh
set -eu
repo=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
build=$(mktemp -d)
trap 'rm -rf "$build"' EXIT HUP INT TERM
cd "$repo"
go build -o "$build/hea" ./cmd/hea
export HEAD_TEST_REPO="$repo"
export HEAD_TEST_BIN="$build/hea"
export HEAD_TEST_PROJECT="$build/project"
export HEAD_TEST_DATA="$build/state"
export PATH="$repo/tests/fixtures:$PATH"
nvim --headless -u NONE --cmd 'lua vim.opt.rtp:append(vim.env.HEAD_TEST_REPO)' -c 'runtime plugin/hea.lua' -c 'lua local ok, err = pcall(dofile, "tests/neovim.lua"); if not ok then print(err); vim.cmd("cquit 1") end' -c 'qa!'
export HEAD_TEST_REOPEN=1
nvim --headless -u NONE --cmd 'lua vim.opt.rtp:append(vim.env.HEAD_TEST_REPO)' -c 'runtime plugin/hea.lua' -c 'lua local ok, err = pcall(dofile, "tests/neovim.lua"); if not ok then print(err); vim.cmd("cquit 1") end' -c 'qa!'
