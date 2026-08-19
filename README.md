# API Testing Suite

Comprehensive API testing toolkit for REST, GraphQL, and gRPC APIs.

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![CI](https://github.com/Qyroxen/api-testing-suite/actions/workflows/ci.yml/badge.svg)](https://github.com/Qyroxen/api-testing-suite/actions/workflows/ci.yml)

> Comprehensive API testing toolkit for REST, GraphQL, and gRPC APIs.

## What is it?

API Testing Suite is a command-line tool built with Go that helps developers comprehensive api testing toolkit for rest, graphql, and grpc apis. It's designed to be fast, reliable, and easy to use.

## Why?

Every developer needs api testing suite — but existing tools are either too complex, too slow, or require cloud dependencies. We built API Testing Suite to be:
- **Fast** — Written in Go for maximum performance
- **Offline** — No cloud dependencies, your data stays on your machine
- **Simple** — Clean CLI interface with sensible defaults
- **Extensible** — Easy to customize and integrate into your workflow

## Features

- **REST, GraphQL, gRPC support** — REST, GraphQL, gRPC support
- **Automated test generation** — Automated test generation
- **Load testing** — Load testing
- **Contract testing** — Contract testing
- **CI/CD integration** — CI/CD integration
- **Report generation** — Report generation

## Quick Start

### Prerequisites

- Go 1.23 or later

### Install

```bash
# Install with go install
go install github.com/Qyroxen/api-testing-suite@latest

# Or build from source
git clone https://github.com/Qyroxen/api-testing-suite.git
cd api-testing-suite
go build -o api-testing-suite .
```

### Usage

```bash
# Basic usage
.api-testing-suite --help

# Example
./api-testing-suite test --url http://localhost:8080/api
```

## Output

```
API Testing Suite v1.0.0

Scanning...

✓ Analysis complete
✓ Results ready

{
  "status": "success",
  "results": [...]
}
```

## Configuration

Create a `.config.yaml` file in your project root:

```yaml
# Configuration options
verbose: true
output: json
timeout: 30s
```

## CLI Flags

```
api testing suite [command]

Flags:
  --path string      Target path (default ".")
  --format string    Output format: json, text (default "text")
  --verbose          Enable verbose output
  --config string    Config file path
  --output string    Output file path
```

## Examples

### Basic Example

```bash
.api-testing-suite --path ./src
```

### Advanced Example

```bash
.api-testing-suite --path ./src --format json --output report.json --verbose
```

### CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run API Testing Suite
  run: |
    go install github.com/Qyroxen/api-testing-suite@latest
    api-testing-suite --path . --format json --output report.json
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [API Reference](docs/api-reference.md)
- [Examples](examples/)
- [Contributing](CONTRIBUTING.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Qyroxen** - [GitHub](https://github.com/Qyroxen)

---

**Found this useful?** Give it a ⭐ on GitHub!
