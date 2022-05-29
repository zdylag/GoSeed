# Go Repository Template

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-success)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/zdylag/GoSeed)](https://github.com/zdylag/GoSeed/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/zdylag/GoSeed.svg)](https://pkg.go.dev/github.com/zdylag/GoSeed)
[![go.mod](https://img.shields.io/github/go-mod/go-version/zdylag/GoSeed)](go.mod)
[![LICENSE](https://img.shields.io/github/license/zdylag/GoSeed)](LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/zdylag/GoSeed/build)](https://github.com/zdylag/GoSeed/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/zdylag/GoSeed)](https://goreportcard.com/report/github.com/zdylag/GoSeed)
[![Codecov](https://codecov.io/gh/zdylag/GoSeed/branch/main/graph/badge.svg)](https://codecov.io/gh/zdylag/GoSeed)

## TEMPLATE INFORMATION

### What is this?

This is a GitHub repository template for Go. It has been created for ease-of-use for anyone who wants to:

- quickly get into Go without losing too much time on environment setup,
- create a new repoisitory with basic Continous Integration.

It includes:

- continous integration via [GitHub Actions](https://github.com/features/actions),
- build automation via [Make](https://www.gnu.org/software/make),
- dependency management using [Go Modules](https://github.com/golang/go/wiki/Modules),
- code formatting using [gofumpt](https://github.com/mvdan/gofumpt),
- linting with [golangci-lint](https://github.com/golangci/golangci-lint),
- unit testing with [testify](https://github.com/stretchr/testify), [race detector](https://blog.golang.org/race-detector), code covarage [HTML report](https://blog.golang.org/cover) and [Codecov report](https://codecov.io/),
- releasing using [GoReleaser](https://github.com/goreleaser/goreleaser),
- dependencies scanning and updating thanks to [Dependabot](https://dependabot.com),
- security code analysis using [CodeQL Action](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning),
- [Visual Studio Code](https://code.visualstudio.com) configuration with [Go](https://code.visualstudio.com/docs/languages/go) and [Remote Container](https://code.visualstudio.com/docs/remote/containers) support.

`Star` this repository if you find it valuable and worth maintaining.

`Watch` this repository to get notified about new releases, issues, etc.

### How do I use it?

1. Sign up on [Codecov](https://codecov.io/) and configure [Codecov GitHub Application](https://github.com/apps/codecov) for all repositories.
1. Click the `Use this template` button (alt. clone or download this repository).
1. In GitHub
   1. Enable `Issues`
   1. Disable `Allow merge commits` and `Allow squash merging`
   1. Enable `Always suggest updating pull request branches`
   1. Enable `Automatically delete head branches`
   1. Enable Tags protection on `*`
   1. Set up the following branch protection rules on `main`:
      - Require a pull request before merging
      - Require approvals
      - Require review from Code Owners
      - Require status checks to pass before merging
      - Require branches to be up to date before merging
      - Status checks that are required
         - Analyze (go)
         - check-commit-message
         - ci (ubuntu-20.04)
         - codecov/patch
         - codecov/project
         - CodeQL
         - release-test
         - test (macos-10.15, 1.16)
         - test (macos-10.15, 1.17)
         - test (macos-10.15, 1.18)
         - test (ubuntu-20.04, 1.16)
         - test (ubuntu-20.04, 1.17)
         - test (ubuntu-20.04, 1.18)
         - test (windows-2019, 1.16)
         - test (windows-2019, 1.17)
         - test (windows-2019, 1.18)
      - Require linear history
   1. Allow all actions and reusable workflows
   1. Require approval for first-time contributors to Fork pull request workflows from outside collaborators
1. In the repository:
   1. Run `make rename`
   1. Replace all occurences of `zdylag/GoSeed` to `your_org/repo_name` in all files.
   1. Delete the following files:
      - [DELETETHIS.go](DELETETHIS.go)
      - [DELETETHIS_test.go](DELETETHIS_test.go)
   1. Update the following files:
      - [CHANGELOG.md](CHANGELOG.md)
      - [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)
      - [CONTRIBUTING.md](CODE_OF_CONDUCT.md)
      - [LICENSE](LICENSE)
      - [README.md](README.md)
   1. Make a pull request to start your new repo off right!

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
