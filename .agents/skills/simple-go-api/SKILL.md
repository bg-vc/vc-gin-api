---
name: simple-go-api
description: Teach, explain, review, or extend the vc-gin-api Gin project while preserving its small Router-to-Handler-to-Service-to-DAO flow. Use when a learner asks how a request works, wants to add a field or endpoint, needs a guided Go backend exercise, or asks Codex to check this repository for overengineering. Do not use to turn the repository into a general framework or microservice platform.
---

# Work with the SIMPLE Go API

## Understand the request

1. Read `AGENTS.md` and the relevant README language.
2. Locate the route, handler, service, DAO, model, and focused tests involved.
3. Explain the current request path in five short steps before proposing a structural change.

## Teach an existing flow

1. Start at the route and follow real function calls; do not explain from generic Gin knowledge alone.
2. Show where HTTP input becomes a model, where business decisions happen, and where SQL executes.
3. Point out one Go concept at a time, using repository code as the example.
4. End with one small exercise the learner can complete in under 30 minutes.

## Extend the API

1. Implement the smallest complete vertical slice through the existing layers.
2. Reuse concrete constructors and existing response/error conventions.
3. Keep SQL parameterized and keep secrets out of logs and tracked files.
4. Avoid new packages and interfaces unless two existing implementations already need them.
5. Keep Redis optional and outside the account/user request path unless the user explicitly changes that rule.
6. Add a focused test for the changed behavior.

## Check simplicity

Reject or simplify a change that introduces speculative layers, generic repositories, generated CRUD, a DI framework, runtime Codex dependencies, or infrastructure unrelated to the learning goal. Explain the concrete cost in files, concepts, or startup steps.

## Verify and hand off

Run `gofmt`, `go vet ./...`, `go test ./...`, and `govulncheck ./...`. Report:

- the request path that changed;
- the behavior a learner can now observe;
- the tests that passed;
- one short next exercise.
