# Bigmile CLI init tool
[![goreleaser](https://github.com/mikkurogue/bigmile-cli/actions/workflows/release.yml/badge.svg)](https://github.com/mikkurogue/bigmile-cli/actions/workflows/release.yml)

A short golang project for bigmile to create a cli tool that can help setup the developers environment with some default cli packages, installing the repo required to start work and potentially installing VSCode or Zed editor for the developer.


## Current support:

Currently the cli supports installing
- eza
- fzf
- bat
- ripgrep
- oh my zsh
- thefuck
- A code editor
  - Zed
  - VSCode
- git-purge alias

Note: to use this, you MUST use the zsh terminal. It makes it easier to develop for if the whole team is using the same shell. Functionally zsh and bash are the same but zsh is slightly newer.

A majority of these commands (should) work for Linux machines too, that use zsh.

Only cli tool that is not recommended for Linux is `thefuck` as this seems to be either super slow or it doesnt work properly on the Linux kernel.

## Planned

 - authenticate with azure repo
 - pull azure repo into target directory
 - Bash support (dynamic options like removing oh my zsh when bash is selected)
 - Check if homebrew is installed - if not then install it first from the script and add the .zshrc options

## How to install

From terminal, run these commands:
`curl -O https://github.com/mikkurogue/bigmile-cli/rel/bigmile.tar.gz`
`curl -O https://github.com/mikkurogue/bigmile-cli/rel/install.sh`

After run the install script:
`chmod +x install.sh`
`./install.sh`

Then once complete you should be able to run the BigMile cli simply by typing `bigmile` to your terminal
