// Package middleware setup middleware fiber
package middleware

import (
	"github.com/attapon-th/go-pkgs/zlog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	flog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

//     "github.com/gofiber/fiber/v2/middleware/compress"
//     "github.com/gofiber/fiber/v2/middleware/cache"

// CORS - Setup Middleware CORS
func CORS() fiber.Handler {
	return cors.New(cors.Config{AllowOrigins: "*"})
}

// LogAccess - Setup Middleware logging request
func LogAccess() fiber.Handler {
	logAPIAccess := zlog.NewConsole()
	if f := viper.GetString("app.logs.access"); f != "" {
		logAPIAccess = zlog.NewLogRollingFile(f)
	}
	return flog.New(flog.Config{
		Format: "[${host}][${latency}][${method}][${status}] ${url}",
		Output: logAPIAccess, //
	})
}

// func Cache() fiber.Handler {
//     return cache.New(cache.Config{
//         Next: func(c *fiber.Ctx) bool {
//             return c.Query("cache") == "false"
//         },
//         Expiration:   30 * time.Minute,
//         CacheControl: true,
//     })
// }

// func Compress() fiber.Handler {
//     return compress.New(compress.Config{
//         Level: compress.LevelBestSpeed,
//     })
// }

// BasicAuth http authentication middleware basicauth
func BasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"user": "pass",
		},
	})
}
