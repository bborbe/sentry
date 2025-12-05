# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.9.0

- update go and deps

## v1.8.5

- update Go version from 1.25.2 to 1.25.3
- update dependencies (Sentry SDK v0.35.3 → v0.36.0, bborbe/run v1.7.8 → v1.8.1, gosec v2.22.9 → v2.22.10)
- remove deprecated golang.org/x/lint dependency
- add comprehensive GoDoc documentation for Client interface and NewClient function
- add sensitive data warning to NewClient documentation
- enhance README with Development, Dependencies, and License sections
- improve error handling with nil checks for eventID in logging
- fix typos in logging messages (orginal → original, execption → exception)
- add defensive nil check for event.Tags initialization
- update copyright headers to include 2025
- update various indirect dependencies (crypto, net, sys, text, etc.)

## v1.8.4

- add golangci-lint configuration file
- enhance Makefile with additional linting and security scanning targets (lint, gosec, trivy, osv-scanner)
- update Go version to 1.25.2
- update dependencies (errors, run, ginkgo)
- add new development tools (golangci-lint, gosec, osv-scanner)
- improve GitHub Actions CI workflow with updated Go version
- streamline README by removing development sections

## v1.8.3

- add EventModifier interface for custom event processing
- add EventModifierFunc type for function-based event modifiers
- add EventModifierList type for sequential modifier application
- add comprehensive test coverage for event modifier functionality
- add detailed godoc documentation for all event modifier types

## v1.8.2

- enhance README.md with comprehensive documentation including status badges, features overview, installation instructions, and usage examples
- add detailed package-level GoDoc comment with usage examples and feature descriptions

## v1.8.1

- add LICENSE file  
- add GitHub workflows for CI and Claude Code integration
- update Go version from 1.24.1 to 1.24.5
- update dependencies (sentry-go, glog, counterfeiter, ginkgo, gomega)
- add golines tool for code formatting
- improve .gitignore patterns

## v1.8.0

- remove vendor
- go mod update

## v1.7.1

- add license
- go mod update 

## v1.7.0

- add data to tags

## v1.6.0

- allow exclude errors

## v1.5.0

- add example
- add tags from context and error

## v1.4.0

- add flush method

## v1.3.0

- simplify newClient

## v1.2.1

- rename proxy roundtripper

## v1.2.0

- add NewClientWithOptions
- add proxy roundTripper

## v1.1.0

- add skip error and report
- go mod update

## v1.0.0

- Initial Version
