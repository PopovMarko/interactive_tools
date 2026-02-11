# Interactive Tools - Pomodoro Timer

[![CI_pomo](https://github.com/PopovMarko/interactive_tools/workflows/CI_pomo/badge.svg)](https://github.com/PopovMarko/interactive_tools/actions/workflows/ci.yaml)
[![Lint](https://github.com/PopovMarko/interactive_tools/workflows/Lint/badge.svg)](https://github.com/PopovMarko/interactive_tools/actions/workflows/lint.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/PopovMarko/interactive_tools)](https://goreportcard.com/report/github.com/PopovMarko/interactive_tools)
<!-- [![codecov](https://codecov.io/gh/PopovMarko/interactive_tools/branch/main/graph/badge.svg)](https://codecov.io/gh/PopovMarko/interactive_tools) -->
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Pomodoro timer written in Go 🍅

## About the Project

This is an interactive Pomodoro timer for productive work using the Pomodoro Technique.

### Pomodoro Technique

- 🍅 25 minutes of work
- ☕ 5 minutes break
- 🎯 4 pomodoros = long break (15-30 minutes)

## Installation

```bash
go install github.com/PopovMarko/interactive_tools/pomo/pomodoro@latest
```

Or clone the repository:

```bash
git clone https://github.com/PopovMarko/interactive_tools.git
cd interactive_tools
go build -o bin/pomodoro ./pomo/pomodoro
```

## Usage

```bash
# Start standard Pomodoro (25 minutes)
./pomodoro

# Start with custom duration
./pomodoro -duration 30m

# Show help
./pomodoro -help
```

## Features

- ⏱️ Customizable timer
- 🔔 Audio notifications
- 📊 Completed pomodoros statistics
- 🎨 Interactive CLI interface

## Development

### Requirements

- Go 1.21 or newer

### Local Development

```bash
# Clone
git clone https://github.com/PopovMarko/interactive_tools.git
cd interactive_tools

# Install dependencies
go mod download

# Run tests
go test -v ./...

# Build
go build -o bin/pomodoro ./pomo/pomodoro

# Run
./bin/pomodoro
```

### Testing

```bash
# Run all tests
go test ./...

# Tests with coverage
go test -cover ./...

# Detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Linting

```bash
# Format code
gofmt -s -w .

# go vet
go vet ./...

# golangci-lint (if installed)
golangci-lint run ./...
```

## CI/CD

The project uses GitHub Actions for automatic checks:

- ✅ Testing on Go 1.21, 1.22, 1.23
- ✅ Linting with golangci-lint
- ✅ Code formatting check
- ✅ Building binaries
- ✅ Test coverage (Codecov)

## Contributing

We welcome your contributions! Please:

1. Fork the project
2. Create a branch for your feature (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Before submitting a PR, make sure that:
- All tests pass ✅
- Code is formatted (`gofmt -s -w .`)
- No linter errors
- Tests added for new functionality

## License

MIT License - see [LICENSE](LICENSE) for details

## Author

**Marko Popov** - [GitHub](https://github.com/PopovMarko)

## Acknowledgments

- Pomodoro Technique created by Francesco Cirillo
- Go community for excellent tools

---

⭐ If you find this project useful - give it a star!
