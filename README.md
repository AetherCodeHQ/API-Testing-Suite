# API Testing Suite

![CI](https://github.com/Qyroxen/API-Testing-Suite/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/API-Testing-Suite?style=social)

> Complete API testing toolkit - automated testing, load testing, and monitoring

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/API-Testing-Suite?style=social)](https://github.com/Qyroxen/API-Testing-Suite/stargazers)

## What is it?

API Testing Suite provides comprehensive API testing including functional tests, load tests, and continuous monitoring.

## Why should you care?

APIs are the backbone of modern apps. This tool ensures they work correctly under any conditions.

## Demo

```bash
./api-test run --url https://api.example.com/users
```

**Output:**
```
API Test Results:
  - 45/45 tests passed
  - Average response time: 120ms
  - Load test: 1000 req/sec sustained
```

## Features

- Functional testing
- Load testing
- Contract testing
- Continuous monitoring
- Report generation

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/API-Testing-Suite.git
cd API-Testing-Suite
go build -o api-test .

# Run
./api-test --url https://api.example.com
```

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--url` | API endpoint | (required) |
| `--method` | HTTP method | `GET` |
| `--load` | Enable load testing | `false` |
| `--concurrent` | Concurrent connections | `10` |
| `--duration` | Load test duration | `60s` |

## Examples

# Functional test
./api-test run --url https://api.example.com/users

# Load test
./api-test run --url https://api.example.com/users --load --concurrent 100

# Contract test
./api-test contract --spec ./openapi.yaml

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/API-Testing-Suite/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/API-Testing-Suite?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/API-Testing-Suite/network/members">
    <img src="https://img.shields.io/github/forks/Qyroxen/API-Testing-Suite?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/API-Testing-Suite/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/API-Testing-Suite" alt="Issues">
  </a>
</p>
