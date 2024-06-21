# Dagger CLI - dagger
[![goreleaser](https://github.com/mikkurogue/bigmile-cli/actions/workflows/release.yml/badge.svg)](https://github.com/mikkurogue/bigmile-cli/actions/workflows/release.yml)
[![](https://dcbadge.limes.pink/api/server/sjuAavPyQt?style=flat)](https://discord.gg/sjuAavPyQt)


A short CLI tool in Go that can help setup the developers environment with some default cli packages, installing VSCode or Zed editor for the developer.

## Installation

`brew tap mikkurogue/mikkurogue`

Once you've added the tap, you may now install the dagger package.

`brew install mikkurogue/mikkurogue/dagger`

note: i know the naming scheme sucks, I'll fix it one day (soon tm)

Once this is installed, you can now simply run the cli in the terminal:
`dagger`.

## Contribution

If you have a feature request, feel free to create a PR. Releases will happen periodically on a whim, contributors can always request a release when a specific feature has been merged.

To get in contact with me, or any contributors, click the Discord badge to join the dagger cli community!
First contributors to contribute and join the discord get a special "founder" role, no extra benefits just a cool little gizmo.

For now as we do not have any real windows support, I do recommend if you are a Windows user to use the WSL2 linux distros to develop and test against this tool. Once Windows support is in a alpha-ish state then WSL2 isnt necessary. Please note, windows support is not on the main radar to create extra development time or priorities for from my end. 

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
 - Bash support (dynamic options like removing oh my zsh when bash is selected)
 - Update .zshrc / .bashrc files for the installed packages
 - Check if homebrew is installed - if not then install it first from the script and add the .zshrc options
 - Windows support
 - pre-checks to see if brew is installed (unix systems)
 - code boilerplate library
   - in `~/dagger/boilerplates/<language>/generated_boilerplate_name.<lang extensoin>` have a set of simple boilerplates like React components, Go main files etc. This idea is still in super infancy



Happy stabbin' (heh get it, cause dagger)!
