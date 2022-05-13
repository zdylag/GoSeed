# Go Repository Template

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/zdylag/TestGoRepo)](https://github.com/zdylag/TestGoRepo/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/zdylag/TestGoRepo.svg)](https://pkg.go.dev/github.com/zdylag/TestGoRepo)
[![go.mod](https://img.shields.io/github/go-mod/go-version/zdylag/TestGoRepo)](go.mod)
[![LICENSE](https://img.shields.io/github/license/zdylag/TestGoRepo)](LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/zdylag/TestGoRepo/build)](https://github.com/zdylag/TestGoRepo/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/zdylag/TestGoRepo)](https://goreportcard.com/report/github.com/zdylag/TestGoRepo)
[![Codecov](https://codecov.io/gh/zdylag/TestGoRepo/branch/main/graph/badge.svg)](https://codecov.io/gh/zdylag/TestGoRepo)

Example use of the go test repository!

## Setup

Below you can find sample instructions on how to set up the development environment.
Of course you can use other tools like [GoLand](https://www.jetbrains.com/go/), [Vim](https://github.com/fatih/vim-go), [Emacs](https://github.com/dominikh/go-mode.el). However take notice that the Visual Studio Go extension is [officially supported](https://blog.golang.org/vscode-go) by the Go team.

### Local Machine

Follow these steps if you are OK installing and using Go on your machine.

1. Install [Go](https://golang.org/doc/install).
1. Install [Visual Studio Code](https://code.visualstudio.com/).
1. Install [Go extension](https://code.visualstudio.com/docs/languages/go).
1. Clone and open this repository.
1. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.

### Development Container

Follow these steps if you do not want to install Go on your machine and you prefer to use a Development Container instead.

1. Install [Visual Studio Code](https://code.visualstudio.com/).
1. Follow [Developing inside a Container - Getting Started](https://code.visualstudio.com/docs/remote/containers#_getting-started).
1. Clone and open this repository.
1. `F1` -> `Remote-Containers: Reopen in Container`.
1. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.

The Development Container configuration mixes [Docker in Docker](https://github.com/microsoft/vscode-dev-containers/tree/master/containers/docker-in-docker) and [Go](https://github.com/microsoft/vscode-dev-containers/tree/master/containers/go) definitions. Thanks to it you can use `go`, `docker`, `docker-compose` inside the container.

## Build

### Terminal

- `make` - execute the build pipeline.
- `make help` - print help for the [Make targets](Makefile).

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)` to execute the build pipeline.

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.

_CAUTION_: Make sure to understand the consequences before you bump the major version. More info: [Go Wiki](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher), [Go Blog](https://blog.golang.org/v2-go-modules).


### How to automate generating git tags for next release version

Auto-tagging can be done in many ways e.g. by using GitHub Actions like:

- [Github Tag Bump](https://github.com/marketplace/actions/github-tag-bump),
- [bumpr](https://github.com/marketplace/actions/bumpr-bump-version-when-merging-pull-request-with-specific-labels),
- [Increment Semantic Version](https://github.com/marketplace/actions/increment-semantic-version),
- [Github Tag](https://github.com/marketplace/actions/github-tag).

However, creating a release tag manually is often the optimal approach. Take notice that this template executes a release workflow each time a git tag with `v` prefix is pushed.

## Contributing

Simply create an issue or a pull request.
