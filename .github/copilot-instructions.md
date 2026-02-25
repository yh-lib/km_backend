# Copilot instructions for km_backend

This file gives concise, actionable guidance for AI coding agents working in this repository.

- **Entry point**: the runnable entry is `clientgo.go` (not `main.go`, which is commented out). Start locally with:

  ```bash
  go run clientgo.go
  # or build: go build -o bin/km_backend ./
  ```

- **Config & env**: runtime configuration uses `github.com/spf13/viper`. Defaults live in [config/config.go](config/config.go#L1-L200).
  - Important env vars: `PORT`, `JWT_SIGN_KEY`, `JWT_EXPIRE_TIME`, `USERNAME`, `PASSWORD` (viper.AutomaticEnv() is used).
  - To change default credentials, edit [config/config.go](config/config.go#L1-L200) or set `USERNAME`/`PASSWORD` environment variables.

- **Routing pattern**:
  - Root router registration: [routers/routers.go](routers/routers.go#L1-L40) registers `/api` and delegates to sub-routers (example: `routers/auth`).
  - Subrouters follow the pattern: `RegisterSubRouter(g *gin.RouterGroup)` and create `group := g.Group("resource")`, then define handlers (see README examples).

- **Authentication & middleware**:
  - JWT middleware is in [middlerwares/middlerwares.go](middlerwares/middlerwares.go#L1-L200). It blocks every request except `/api/auth/login` and `/api/auth/logout`.
  - Middleware expects an `Authorization` header containing the token string and uses `utils/jwtutils` to parse it.
  - Note: the middleware writes a JSON status response (via `config.NewRetrunData()`) before calling `c.Next()` — controllers may also write JSON. Be careful to avoid writing conflicting responses.

- **JWT helpers**: token creation and parsing are in [utils/jwtutils/jwtutils.go](utils/jwtutils/jwtutils.go#L1-L200). It uses `config.JwtSignKey` and `config.JwtExpTime`.

- **Logging**: use the logs wrapper `utils/logs/logs.go` (functions `Info`, `Debug`, `Error`, etc.) so output follows the project's JSON format and caller formatting. See [README.md](README.md#L1-L200) for examples.

- **Common data shape**: controllers return `config.ReturnData` (constructed with `config.NewRetrunData()`) — keep to this structure `{status, message, data}`.

- **Dependencies**: primary frameworks are `github.com/gin-gonic/gin` for HTTP and `github.com/golang-jwt/jwt/v5` for JWT (see `go.mod`).

- **Where to add features**:
  - New API groups: add a folder under `routers/` and `controllers/`, implement `RegisterSubRouter` and the corresponding controller functions.
  - Use `config.NewRetrunData()` as the response template and `logs.Info(...)` for logging.

- **Developer notes / gotchas discovered from code**:
  - `main.go` contains an alternate (commented) entry; `clientgo.go` is the actual runnable main.
  - Middleware writes a response early (200/401 JSON) and still sets `claims` in context; controllers should read `c.Get("claims")` if they need user info.
  - Defaults use MD5-hashed strings in `config.config.go` for username/password; when testing, set env vars to override.

- **Quick edits & testing**:
  - To test login flow: POST to `/api/auth/login` then include the returned token in `Authorization` header for other requests.
  - To change log level or port quickly, set `LOG_LEVEL` and `PORT` env vars before running.

If anything above is unclear or you want examples for a specific change (new router, new controller, or CI step), tell me which area to expand.
