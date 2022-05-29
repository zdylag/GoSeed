# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/zdylag/GoSeed/compare/...HEAD)

## [0.1.0](https://github.com/zdylag/GoSeed/releases/tag/v0.1.0) - 2022-05-28

### Added

- Backbone for template project based on [golang-templates/seed](https://github.com/golang-templates/seed)
- GitHub issue templates for bugs, feature requests, and tasks
- GitHub setup instructions to README
- Commit format linter
- Additional useful extensions to the devcontainers starting list
- Contributing guidelines

### Changed

- Go version for the devcontainer to 1.18
- Go version for the default library to 1.16
- Full CI to only run on a single Go version on a single OS (Ubuntu 20.04)
- CI tests now run on multiple Go versions in addition to Operating Systems
- Dependabot sends commits formated according to the style of the repo
- Default project files from `main` to `DELETETHIS`
- Default license from Unilicense to MIT

### Removed

- Binary output support
- Installation instructions