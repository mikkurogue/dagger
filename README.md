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

Requirements: Github access token

Go to https://github.com/settings/tokens/new and create a new token. I recommend creating a token that does not expire, but its up to you.

In your terminal, create the api access token env variable
`export HOMEBREW_GITHUB_API_TOKEN=<the token you just generated>`

Once you have created the HOMEBREW_GITHUB_API_TOKEN environment variable, you have to add the brew tap to the package.

`brew tap mikkurogue/mikkurogue`

Once you've added the tap, you may now install the bigmile-cli package

`brew install mikkurogue/mikkurogue/bigmile-cli`

note: i know the naming scheme sucks, I'll fix it one day (soon tm)

If you are feeling brave, you can also run the command

`brew install bigmile-cli`

However this does not always guarantee you go to the right brew tap, but if you're feeling frisky go right ahead!


Once this is installed, you can now simply run the cli in the terminal:
`bigmile-cli`

Happy BigMile cli-ing