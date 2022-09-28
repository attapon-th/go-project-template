# Fiber v2

- [Github Fiber v2](https://github.com/gofiber/fiber)

- [Docs Fiber v2](https://docs.gofiber.io/)

## List Packages

| Name(link)                                                   | Import                                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ------------------------------------------------------------ | ---------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [lancet](https://github.com/duke-git/lancet)                 |                                          | [algorithm sort, search](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md) \| [cryptor](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md) \| [datetime](https://github.com/duke-git/lancet/blob/main/docs/datetime.md) \| [fileutil](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md) \| [maputil](https://github.com/duke-git/lancet/blob/main/docs/maputil.md) \| [netutil](https://github.com/duke-git/lancet/blob/main/docs/netutil.md) \| [random](https://github.com/duke-git/lancet/blob/main/docs/random.md) \| [slice](https://github.com/duke-git/lancet/blob/main/docs/slice.md) \| [strutil](https://github.com/duke-git/lancet/blob/main/docs/strutil.md) |
| [helper]()                                                   | `"github.com/attapon-th/go-pkgs/helper"` | helper function                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [zlog](https://github.com/attapon-th/go-pkgs/tree/main/zlog) | `"github.com/attapon-th/go-pkgs/zlog"`   | logger                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| [task](https://github.com/attapon-th/go-pkgs/tree/main/task) | `"github.com/attapon-th/go-pkgs/task"`   | a goroutine pool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| [jsoniter](https://github.com/json-iterator/go)              | `jsoniter "github.com/json-iterator/go"` | json                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [cache](https://github.com/patrickmn/go-cache)               | `"github.com/patrickmn/go-cache"`        | in memory cache                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [gorm](https://gorm.io/docs/)                                | `"gorm.io/gorm"`                         | ORM library                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [sqlx](https://jmoiron.github.io/sqlx/)                      | `"github.com/jmoiron/sqlx"`              | `sql` package same `database/sql`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| [viper](https://github.com/spf13/viper)                      | `"github.com/spf13/viper"`               | Configuration                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |

## Directory Layouts

- cmd

- internal
  
  - controller
  - middleware
  
  - model

- pkg -- global package

- .config.yaml -- _default config_

- main.go  -- _auto create with_ `cobra-cli init --viper --author "Firstname Lastname"`

## Initial Project

```shell
git init
git remote add origin https://github.com/_USER_/_REPO_.git # please change _USER_, _REPO_
go mod init github.com/_USER_/_REPO_ # please change _USER_, _REPO_
```

### Use `cobra-cli` and load config with `viper`

#### [cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md) - Cobra Generator Documentation

```shell
go install github.com/spf13/cobra-cli@latest # install cobra-cli

## -Create Direcotory- ##
mkdir -p internal/controller \
&& mkdir -p internal/middleware \
&& mkdir -p internal/model \
&& mkdir -p pkg

## -add command cli- ##
cobra-cli init --viper --author "Firstname Lastname"

cobra-cli add serv

go mod tidy
## - test -
go run main.go --help
```

## Configs `./.config.yaml`

```yaml
version: 1.0.0
app:
  dev: true
  listen: localhost # production set: 0.0.0.0
  prefix: /api/v1
  port: 8888
  cpu: 2 # if fiber.Config.Prefork=True
  logs:
    log: console
    access: storage/logs/access.log
```

---

## Fiber App server `./internal/server.go`

```go
// Package internal usecase for project
package internal
import (
    jsoniter "github.com/json-iterator/go"
    "github.com/gofiber/fiber/v2"
)

var (
    Version     string = "1.0.0"
    Build       string = ""
    Timestamp   string = ""
)

func init(){
    app := NewFiber()
    controller.EndpointPing(app, viper.GetString("app.prefix"))
}

func NewFiber() *fiber.App{
    app := fiber.New  
    config := fiber.Config{
        AppName:              "fiber-api",
        Prefork:               false,
        StrictRouting:         true,
        CaseSensitive:         true,
        Immutable:             false,
        BodyLimit:             4 * 1024 * 1024, // 4mb body size
        Concurrency:           256 * 1024, 
        JSONEncoder:           jsoniter.Marshal,
        JSONDecoder:           jsoniter.Unmarshal,
        ErrorHandler:          ErrorHandlerResponseJSON,
    }
    if nc := viper.GetInt("app.cpu"); nc == 1 || nc < 0 {
        config.Prefork = false
    } else if nc == 0 {
        config.Prefork = true
    } else if nc > 1 {
        config.Prefork = true
        runtime.GOMAXPROCS(nc)
    }
    app := fiber.New(config)
    return app
}


func ErrorHandlerResponseJSON(ctx *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError
    msgError := "Error!!!"
    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
    }
    if err != nil {
        msgError = err.Error()mi
    }
    resError := fiber.Map{
        "msg": msgError,
        "ok":  false,
    }
    return ctx.Status(code).JSON(resError)
}
```

---

## Controller

### Ping `./internal/controller/pingController.go`

```go
package controller

import (
    "github.com/gofiber/fiber/v2"
)
func EndpointPing(rt fiber.Router, prefixUrl string){
    r := rt.Group(prefixUrl, middleware.CORS(), middleware.LogAccess() )

    r.Get("/ping", PingHandler)
}

func PingHandler(c *fiber.Ctx) error{
    return fiber.Mab{ok: true, msg: "ping succcess"}
}
```

---

## Middleware `./internal/middleware/middleware.go`

```go
// Package middleware setup middleware fiber 
package middleware
// import(
    // "github.com/gofiber/fiber/v2/middleware/compress"
    // "github.com/gofiber/fiber/v2/middleware/cors"
    // "github.com/attapon-th/go-pkgs/zlog"
    // flog "github.com/gofiber/fiber/v2/middleware/logger"
    // "github.com/gofiber/fiber/v2/middleware/cache"
    // "github.com/gofiber/fiber/v2/middleware/basicauth"
// )

func CORS() fiber.Handler {
    return cors.New(cors.Config{AllowOrigins: "*"})
}

func LogAccess() fiber.Handler {
    logApiAcccess := zlog.NewLogRollingFile(viper.GetString("app.logs.access")
    return flog.New(flog.Config{
        Format: "[${host}][${latency}][${method}][${status}] ${url}",
        Output: logApiAcccess, //
    })
}

func Cache() fiber.Handler {
    return cache.New(cache.Config{
        Next: func(c *fiber.Ctx) bool {
            return c.Query("cache") == "false"
        },
        Expiration:   30 * time.Minute,
        CacheControl: true,
    })
}

func Compress() fiber.Handler {
    return compress.New(compress.Config{
        Level: compress.LevelBestSpeed,
    })
}


func BasicAuth() fiber.Handler {
    return basicauth.New(basicauth.Config{
        Users: map[string]string{
            "user":  "pass",
        },
    })
}
```

---
