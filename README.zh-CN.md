# vc-gin-api

[English](README.md) | [简体中文](README.zh-CN.md)

**SIMPLE —— 给想学 Go、转 Go 的人，一个看得懂、跑得起来、改得动的 Gin 实战项目。**

LLM 写 Go 的能力越来越强，问题是它们经常写得太多。

`vc-gin-api` 创建于 2019 年，最初用于帮助后端开发者通过一个真实、易读的项目快速上手 Go。七年后，Codex 帮助它完成了克制的现代化，而没有把它变成框架。它具有真实后端项目的基本形状：配置、认证、MySQL、测试和优雅退出；但它有意不宣称自己是可直接上线的生产系统。

## 一个请求，五个看得见的步骤

```text
HTTP -> Router -> Handler -> Service -> DAO -> MySQL
```

例如，新增用户只有一条直线调用链：

```text
POST /api/user/auth/add
  -> router/router.go
  -> handler/user_handler.go
  -> service/user_service.go
  -> dao/user_dao.go
  -> MySQL
```

没有 DI 容器，没有通用 Repository，也没有代码生成器。

## 三步运行

需要 Go 1.26.5+，以及带 Compose 的 Docker。

```bash
cp .env.example .env
docker compose up -d --wait
go run .
```

检查服务：

```bash
curl http://127.0.0.1:8080/api/check/health
```

本地数据库会初始化一个仅用于演示的账号：

```text
用户名：admin
密码：simple-go-demo
```

不要在本项目之外复用这些本地演示值。

## 跑通完整 API 流程

### 1. 登录

```bash
curl -X POST http://127.0.0.1:8080/api/account/login \
  -H 'Content-Type: application/json' \
  -d '{"name":"admin","password":"simple-go-demo"}'
```

从响应中复制 `data.token`：

```bash
TOKEN='在这里粘贴-token'
```

### 2. 查询用户

```bash
curl 'http://127.0.0.1:8080/api/user/common/queryAll?start=0&limit=20'
```

可选筛选参数为 `name` 和 `address`。

### 3. 新增用户

```bash
curl -X POST http://127.0.0.1:8080/api/user/auth/add \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"Alice","age":28,"address":"Shanghai"}'
```

### 4. 修改这个用户

使用新增接口返回的用户 `id`：

```bash
curl -X POST http://127.0.0.1:8080/api/user/auth/update \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"id":2,"name":"Alice","age":29,"address":"Beijing"}'
```

### 5. 修改账号密码

```bash
curl -X POST http://127.0.0.1:8080/api/account/updatePass \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"admin","oldPwd":"simple-go-demo","newPwd":"my-new-password"}'
```

所有接口使用同一种响应结构：

```json
{"code":0,"msg":"ok","data":{}}
```

HTTP 状态码仍然表达真实结果：输入错误为 `400`、认证失败为 `401`、数据不存在为 `404`、未预期的服务错误为 `500`。

## 路由

| 方法 | 路径 | 需要认证 |
| --- | --- | --- |
| GET | `/api/check/health` | 否 |
| POST | `/api/account/login` | 否 |
| POST | `/api/account/updatePass` | 是 |
| GET | `/api/user/common/queryAll` | 否 |
| POST | `/api/user/auth/add` | 是 |
| POST | `/api/user/auth/update` | 是 |

## 项目目录

```text
main.go      加载配置、连接依赖、启动和关闭 HTTP 服务
router/      匹配 URL 并应用中间件
handler/     绑定 HTTP 输入并发送 HTTP 响应
service/     处理账号和用户业务判断
dao/         执行参数化 SQL
model/       定义请求、响应和数据库数据
pkg/         小型基础设施辅助：配置、JWT、MySQL、Redis
```

依赖通过 `NewUserService` 这样的简单构造函数传入。项目没有全局数据库客户端，也没有依赖注入框架。

## 配置

应用只在启动时读取一次 `.env`，并且 `.env` 已被 Git 忽略。

| 变量 | 用途 | 默认值 |
| --- | --- | --- |
| `VC_HTTP_ADDRESS` | HTTP 监听地址 | `:8080` |
| `VC_GIN_MODE` | Gin 运行模式 | `debug` |
| `VC_JWT_SECRET` | JWT HMAC 密钥，至少 32 个字符 | 必填 |
| `VC_JWT_EXPIRY` | Token 有效时间 | `30m` |
| `VC_DB_HOST`、`VC_DB_PORT` | MySQL 地址 | `127.0.0.1:3306` |
| `VC_DB_NAME`、`VC_DB_USER`、`VC_DB_PASSWORD` | MySQL 数据库配置 | 密码必填 |
| `VC_REDIS_ENABLED` | 启用可选 Redis 示例 | `false` |

### 可选 Redis 示例

账号和用户 API 不依赖 Redis。要学习小型 JSON 缓存辅助函数：

1. 在 `.env` 中设置 `VC_REDIS_ENABLED=true`。
2. 运行 `docker compose --profile redis up -d --wait`。
3. 重新运行 `go run .`，阅读 `pkg/redis.go`。

如果本机已有服务占用 `3306` 或 `6379`，请在 `.env` 中修改 `VC_DB_PORT`，或者同时修改 `VC_REDIS_PORT` 和 `VC_REDIS_ADDRESS`。

## 使用 Codex 学习

仓库内置了一个遵循 [Codex 官方 Skills 格式](https://learn.chatgpt.com/docs/build-skills)的教学 Skill：

```text
.agents/skills/simple-go-api/SKILL.md
```

在 Codex 中使用 `$simple-go-api` 调用。推荐从这些提示开始：

```text
Use $simple-go-api to explain the login request from Router to MySQL.
Use $simple-go-api to add an email field without adding a new abstraction.
Use $simple-go-api to review this change for overengineering.
```

这个 Skill 用于讲解和检查现有调用链，不是运行时依赖，也不会生成新的框架。

## 验证改动

```bash
gofmt -w main.go dao handler model pkg router service
go vet ./...
go test ./...
govulncheck ./...
```

测试覆盖 JWT 过期和校验、bcrypt 登录、响应状态映射、缺少认证、参数化用户操作，以及 Redis 开启和关闭两种行为。

## 明确不做什么

- 微服务、Clean Architecture 或通用后端框架
- CRUD 生成、OpenAPI 生成或 DI 容器
- Kubernetes、镜像仓库、可观测性套件或云部署教程
- 完整权限系统、限流器或生产安全承诺
- 在账号或用户核心调用链中强制使用 Redis

真正需要这些能力时再增加，而不是提前增加。

## License

[MIT](LICENSE)
