# vimgrep
A small wrapper for grep that looks for strings in files with the supplied extension
and provides a list of options. All the selected files will then be opened in vim.

Requires the vim extension https://github.com/wsdjeg/vim-fetch to work properly

example: vimgrep go At least

Will find the phrase At least recursivly from the current folder. Case sensitive.
