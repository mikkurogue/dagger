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

  Note: to use this, you MUST use the zsh terminal. It makes it easier to develop for if the whole team is using the same shell. Functionally zsh and bash are the same but zsh is slightly newer.

  ## Planned

  - authenticate with azure repo
  - pull azure repo into target directory
  - git purge alias
