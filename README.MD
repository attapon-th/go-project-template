# Taskfile

## Install taskfile: https://taskfile.dev/installation/

```shell
go install github.com/go-task/task/v3/cmd/task@latest
```

## Use get [Taskfile.yaml](https://raw.githubusercontent.com/attapon-th/go-project-template/main/Taskfile.yaml) into root directory project

```shell
curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/Taskfile.yaml > Taskfile.yaml \
&& task goinit \
&& cobra-cli init --viper \
&& echo "/vendor" >> .gitignore
```



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

---

# [Fiber v2 API](https://docs.gofiber.io/)

## [Fiber v2 Layouts](https://github.com/attapon-th/go-project-template/tree/main/tlp/fiberv2)

> cmd   -- create by `cobra-cli`
> 
> > root.go 
> > 
> > serv.go
> 
> internal
> 
> > controller
> > 
> > middleware
> > 
> > model
> 
> pkg  -- global package
> 
> .confg.yaml
> 
> main.go -- create by `cobra-cli`



## Initialize Project Fiber v2

```shell
# cd to project
## create go project
curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/Taskfile.yaml > Taskfile.yaml \
&& task goinit \
&& cobra-cli init --viper \
&& echo "/vendor" >> .gitignore
## end create go project 

## create fiber
mkdir -p internal/controller \
&& mkdir -p internal/middleware \
&& mkdir -p internal/model \
&& mkdir -p pkg \
&& curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/.config.yaml > .config.yaml \
&& curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/internal/fiber.go > internal/fiber.go \
&& curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/internal/controller/pingController.go > internal/controller/pingController.go \
&& curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/internal/middleware/middleware.go > internal/middleware/middleware.go
&& cobra-cli add serv

## Get swagger if your use https://github.com/swaggo/swag
curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/_main-dev.go > _main-dev.go \
&& curl https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/_main-prd.go > _main-prd.go


```

### Edit file

- `cmd/serv.go` -- for run server filber call in `internal.Serv()`

- internal/fiber.go -- edit call any `controller.Endpoint{ANY}`  for setup route

### 

## Configs `./.config.yaml`

> Pathfile: `./.config.yaml`
> 
> Example: [.config.yaml](https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/.config.yaml) 

## Setup Fiber

> Pathfile: `./internal/fiber.go`
> 
> Example: [fiber.go](https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/internal/fiber.go)

---

## Controller

> Pathfile: `./internal/controller/pingController.go`
> 
> Example: [pingController.go](https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/internal/controller/pingController.go)

## Middleware `./internal/middleware/middleware.go`

> Pathfile: `./internal/middleware/middleware.go`
> 
> Example: [middleware.go](https://raw.githubusercontent.com/attapon-th/go-project-template/main/tlp/fiberv2/internal/middleware/middleware.go)
