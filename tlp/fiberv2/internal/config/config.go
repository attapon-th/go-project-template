// Package config in project
package config

import (
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal"
	"github.com/spf13/viper"
)

// version: 1.0.0
// app:
//   dev: false
//   listen: 0.0.0.0
//   prefix: /api/v1
//   port: 8888
//   cpu: 2 # if cpu < 2: fiber.Config.Prefork=false
//   logs:
//     log: console
//     access: storage/logs/access.log

// DefualtConfig in project
func SetDefualtConfigs() {
	v := viper.GetViper()
	v.SetDefault("version", internal.Version)
	v.SetDefault("app.dev", true)
	v.SetDefault("app.listen", "0.0.0.0")
	v.SetDefault("app.port", 80)
	v.SetDefault("app.prefix", "/api/v1")
	v.SetDefault("app.cpu", 1)
	v.SetDefault("app.log", "console")
	v.SetDefault("app.access", "storage/logs/access.log")
}
