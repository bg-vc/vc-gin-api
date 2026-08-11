# vc-gin-api agent guide

Read `README.md` before changing this repository.

## Product rule: SIMPLE

This repository teaches people who are learning Go or moving to Go. Keep the request path visible:

```text
Router -> Handler -> Service -> DAO -> MySQL
```

- Prefer changing an existing file over creating a new package.
- Do not add an interface, framework, generator, or abstraction until real duplication requires it.
- Prefer the Go standard library when it keeps the code easier to read.
- Keep handlers responsible for HTTP, services responsible for decisions, and DAOs responsible for SQL.
- Pass dependencies through small constructors. Do not add global mutable state or a DI container.
- Keep Redis optional. The account and user APIs must work without it.
- Never log passwords, tokens, database DSNs, or secret configuration values.
- Preserve the six documented API paths unless a task explicitly requests a breaking change.

## Change workflow

1. Trace the current request from `router/` through `handler/`, `service/`, and `dao/`.
2. Make the smallest complete vertical change.
3. Add or update a focused test next to the changed behavior.
4. Run:

```bash
gofmt -w main.go dao handler model pkg router service
go vet ./...
go test ./...
govulncheck ./...
```

5. In the handoff, explain the request path and why no extra abstraction was needed.
