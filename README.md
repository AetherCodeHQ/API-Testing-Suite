# API Testing Suite

![CI](https://github.com/Qyroxen/API-Testing-Suite/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/API-Testing-Suite/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/API-Testing-Suite?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/API-Testing-Suite)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/API-Testing-Suite)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/API-Testing-Suite?style=social)](https://github.com/Qyroxen/API-Testing-Suite/stargazers)

## What is it?

API Testing Suite is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/API-Testing-Suite.git
cd API-Testing-Suite
go build -o apitestingsuite .

# Run
./apitestingsuite --help
```

## CLI Usage

```bash
# Basic usage
./apitestingsuite

# With flags
./apitestingsuite --verbose --output json

# Get help
./apitestingsuite --help
```

## Examples

```bash
# Example 1
./apitestingsuite example1

# Example 2
./apitestingsuite example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o apitestingsuite .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/API-Testing-Suite/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/API-Testing-Suite?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/API-Testing-Suite/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/API-Testing-Suite?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/API-Testing-Suite/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/API-Testing-Suite" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/API-Testing-Suite/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/API-Testing-Suite" alt="Pull Requests">
  </a>
</p>
