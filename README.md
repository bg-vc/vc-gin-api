# vc-gin-api

[English](README.md) | [简体中文](README.zh-CN.md)

**SIMPLE — A small, practical Gin API for people learning or moving to Go.**

LLMs can write more Go than ever. The problem is they often write too much.

`vc-gin-api` was created in 2019 to help backend developers get started with Go through a real, readable project. Seven years later, Codex helped modernize it without turning it into a framework. It is production-shaped—configuration, authentication, MySQL, tests, and graceful shutdown are here—but it is intentionally not production-ready software.

## One request, five visible steps

```text
HTTP -> Router -> Handler -> Service -> DAO -> MySQL
```

For example, adding a user follows one straight path:

```text
POST /api/user/auth/add
  -> router/router.go
  -> handler/user_handler.go
  -> service/user_service.go
  -> dao/user_dao.go
  -> MySQL
```

No DI container. No generic repository. No code generator.

## Run in three steps

Requirements: Go 1.26.5+ and Docker with Compose.

```bash
cp .env.example .env
docker compose up -d --wait
go run .
```

Check the server:

```bash
curl http://127.0.0.1:8080/api/check/health
```

The local database is initialized with one demo-only account:

```text
username: admin
password: simple-go-demo
```

Do not reuse these local demonstration values outside this project.

## Try the complete API flow

### 1. Log in

```bash
curl -X POST http://127.0.0.1:8080/api/account/login \
  -H 'Content-Type: application/json' \
  -d '{"name":"admin","password":"simple-go-demo"}'
```

Copy `data.token` from the response:

```bash
TOKEN='paste-token-here'
```

### 2. Query users

```bash
curl 'http://127.0.0.1:8080/api/user/common/queryAll?start=0&limit=20'
```

Optional filters are `name` and `address`.

### 3. Add a user

```bash
curl -X POST http://127.0.0.1:8080/api/user/auth/add \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"Alice","age":28,"address":"Shanghai"}'
```

### 4. Update that user

Use the returned user `id`:

```bash
curl -X POST http://127.0.0.1:8080/api/user/auth/update \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"id":2,"name":"Alice","age":29,"address":"Beijing"}'
```

### 5. Change the account password

```bash
curl -X POST http://127.0.0.1:8080/api/account/updatePass \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"admin","oldPwd":"simple-go-demo","newPwd":"my-new-password"}'
```

Every endpoint uses the same response shape:

```json
{"code":0,"msg":"ok","data":{}}
```

HTTP status codes still describe success or failure: `400` for invalid input, `401` for authentication failures, `404` for missing data, and `500` for unexpected server errors.

## Routes

| Method | Path | Authentication |
| --- | --- | --- |
| GET | `/api/check/health` | No |
| POST | `/api/account/login` | No |
| POST | `/api/account/updatePass` | Yes |
| GET | `/api/user/common/queryAll` | No |
| POST | `/api/user/auth/add` | Yes |
| POST | `/api/user/auth/update` | Yes |

## Project layout

```text
main.go      Load config, connect dependencies, start and stop HTTP
router/      Match URLs and apply middleware
handler/     Bind HTTP input and send HTTP responses
service/     Make account and user decisions
dao/         Execute parameterized SQL
model/       Define request, response, and database data
pkg/         Small infrastructure helpers: config, JWT, MySQL, Redis
```

Dependencies are passed through small constructors such as `NewUserService`. There are no global database clients and no dependency-injection framework.

## Configuration

The app reads `.env` once at startup. `.env` is ignored by Git.

| Variable | Purpose | Default |
| --- | --- | --- |
| `VC_HTTP_ADDRESS` | HTTP listen address | `:8080` |
| `VC_GIN_MODE` | Gin mode | `debug` |
| `VC_JWT_SECRET` | JWT HMAC secret, at least 32 characters | Required |
| `VC_JWT_EXPIRY` | Token lifetime | `30m` |
| `VC_DB_HOST`, `VC_DB_PORT` | MySQL address | `127.0.0.1:3306` |
| `VC_DB_NAME`, `VC_DB_USER`, `VC_DB_PASSWORD` | MySQL database credentials | Password required |
| `VC_REDIS_ENABLED` | Enable the optional Redis example | `false` |

### Optional Redis example

The account and user APIs do not depend on Redis. To explore the small JSON cache helper:

1. Set `VC_REDIS_ENABLED=true` in `.env`.
2. Run `docker compose --profile redis up -d --wait`.
3. Restart `go run .` and read `pkg/redis.go`.

If a local service already uses `3306` or `6379`, change `VC_DB_PORT`, or both `VC_REDIS_PORT` and `VC_REDIS_ADDRESS`, in `.env`.

## Learn with Codex

This repository includes a focused Codex skill following the [official Skills format](https://learn.chatgpt.com/docs/build-skills):

```text
.agents/skills/simple-go-api/SKILL.md
```

Invoke it from Codex with `$simple-go-api`. Good first prompts:

```text
Use $simple-go-api to explain the login request from Router to MySQL.
Use $simple-go-api to add an email field without adding a new abstraction.
Use $simple-go-api to review this change for overengineering.
```

The skill teaches and checks the existing flow. It is not a runtime dependency and does not generate a new framework.

## Verify a change

```bash
gofmt -w main.go dao handler model pkg router service
go vet ./...
go test ./...
govulncheck ./...
```

Tests cover JWT expiry and validation, bcrypt login, response status mapping, missing authentication, parameterized user operations, and Redis enabled/disabled behavior.

## Deliberate non-goals

- Microservices, Clean Architecture, or a generic backend framework
- Generated CRUD, OpenAPI generation, or a DI container
- Kubernetes, a container registry, observability stacks, or cloud deployment recipes
- A complete authorization system, rate limiter, or production security guarantee
- Redis in the core account or user request path

If you need those things, add them when your application actually needs them—not before.

## License

[MIT](LICENSE)
