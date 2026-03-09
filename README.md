# zen

> **Less is more. Clear is better than clever.**

A minimal Go CLI application skeleton — designed to be a clean starting point, not a framework.

## Philosophy

This project embraces Go's core design principles:

- **Less is more** — Minimal dependencies, no unnecessary abstractions
- **Clear is better than clever** — Simple, readable code over clever tricks
- **No forced architecture** — Use DDD, MVC, or whatever fits your needs
- **Gopher to the core** — Idiomatic Go, standard library first

## Features

Built-in essentials for production-ready CLI applications:

- **Cobra Integration** — Extended command framework with organized flag management
- **Configuration Management** — Viper-based config with file, env var, and flag support
- **Structured Logging** — Standard library `log/slog` with JSON/text formats
- **Graceful Shutdown** — Signal-aware context for clean termination (SIGTERM/SIGINT)
- **Organized Flag Sets** — Grouped flags for better help output

## What It Is Not

- ❌ A framework prescribing how you should architect your app
- ❌ A DDD/Clean Architecture template forcing patterns on you
- ❌ A dependency-heavy starter with dozens of third-party packages
- ❌ A "batteries included" solution you need to strip down

## Project Structure

```text
zen/
├── cmd/              # CLI commands (Cobra)
├── internal/         # Private application code
├── pkg/
│   ├── app/          # Application framework (Cobra extensions, config, flags)
│   ├── signal/       # Graceful shutdown with signal-aware context
│   └── logs/         # Structured logging wrapper
└── main.go           # Entry point
```

The structure is intentionally minimal. Adapt it to your needs.

## Usage

```bash
# Build
go build -o zen

# Run
./zen
```

## Acknowledgments

Inspired by and patterns extracted from [Kubernetes](https://github.com/kubernetes/kubernetes) — signal handling, flag management, and Cobra usage patterns, stripped down to the essentials for a clean CLI application foundation.

## License

MIT © 2026 tsukiyo
