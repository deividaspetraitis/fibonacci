# About

serverd is HTTP Server interface implementation of Fibonacci sequence counter.

# Usage

Please run program with `--help` flag to see available configuration options if running manually.

# Build and run with docker

Copy `.env.example` to `.env` to the project root and update configuration values as necessary.

Run application:

```bash
docker compose up
```

# Test

```bash
curl 'http://localhost/current' -v
```

See [README.md](../../README.md) for available endpoints.
