# Bigmile CLI init tool

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

## How to use? (future, not working yet)

When installed, just run `bigmile init` and you should see the cli tool open.
