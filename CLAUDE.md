# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

cLive is a Go CLI tool that automates terminal operations and displays them live in a web browser. It's designed for creating terminal demos, tutorials, and automated workflows using YAML configuration files.

## Essential Commands

### Development
- `go run .` - Run the application directly from source
- `go build` - Build the binary
- `go test ./... -race -coverprofile=coverage.out -covermode=atomic` - Run tests (CI command)
- `go test ./internal/config -v` - Run specific package tests with verbose output

### Testing Specific Packages
- `go test ./internal/config` - Test configuration package
- `go test ./internal/ttyd` - Test ttyd integration
- `go test ./cmd` - Test CLI commands

### Linting and Quality
- `golangci-lint run --verbose ./...` - Run linter (exact CI command)
- `mise install` - Install development tools (Go 1.24.3, golangci-lint 2.1.6)

### Application Usage
- `./clive init` - Initialize a new clive.yml config file
- `./clive start` - Start a cLive session
- `./clive validate` - Validate configuration file
- `./clive completion <shell>` - Generate shell completions

## Architecture

### Core Components
- **`cmd/`**: CLI commands built with Cobra framework
  - `init.go` - Creates default configuration
  - `start.go` - Main TUI application using Bubble Tea
  - `validate.go` - Config validation
  - `notify.go` - Notifications (likely for completion)

- **`internal/config/`**: Configuration system with YAML parsing and JSON schema validation
  - Supports 6 action types: TypeAction, KeyAction, SleepAction, PauseAction, CtrlAction, ScreenshotAction
  - Settings for terminal appearance, browser options, and behavior

- **`internal/ui/`**: Bubble Tea TUI framework integration
  - Manages application state and user interactions during playback

- **`internal/ttyd/`**: Integration with external ttyd process
  - Provides web-based terminal interface that displays in browser

- **`internal/browser/`**: Browser automation using go-rod
  - Launches and controls browser to display the web terminal

- **`internal/util/`**: Common utilities for strings, slices, integers, versions

### Key Dependencies
- **Cobra**: CLI framework
- **Bubble Tea**: TUI framework for the interactive interface
- **go-rod**: Browser automation
- **ttyd**: External dependency for web terminal (version 1.7.4+)

## Development Workflow

1. **Prerequisites**: Ensure ttyd (1.7.4+) is installed via `brew install ttyd`
2. **Tool Management**: Uses mise for Go and tool versioning
   - `mise install` - Install Go 1.24.3, golangci-lint 2.1.6, goreleaser 2.9.0
   - Tool versions specified in `mise.toml`
3. **Testing**: Uses testify framework with table-driven tests and custom mocks
4. **Schema Validation**: `schema.json` provides IDE autocompletion for `clive.yml` files
5. **CI/CD**: GitHub Actions with 10-minute timeouts, codecov integration, release-please automation

## Configuration

The application uses `clive.yml` files with two main sections:
- **settings**: Terminal appearance, browser options, timing defaults
- **actions**: Sequence of terminal actions to execute (type, key, ctrl, sleep, pause, screenshot)

JSON schema validation ensures configuration correctness and provides IDE autocompletion.

## Testing Notes

- Uses testify framework for assertions and mocking
- Custom mocks in `internal/net/mocks_test.go` for network components
- Table-driven test patterns (see `util/int_test.go`)
- `TestMain` functions for setup/teardown in test suites
- Race detection enabled in CI pipeline
- Configuration parsing has comprehensive test coverage including edge cases

## Release Process

Uses goreleaser for cross-platform binary distribution with multiple installation methods (Homebrew, go install, direct download).