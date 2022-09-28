// Package middleware setup middleware fiber
package middleware

// import(
//     "github.com/gofiber/fiber/v2/middleware/compress"
//     "github.com/gofiber/fiber/v2/middleware/cors"
//     "github.com/attapon-th/go-pkgs/zlog"
//     flog "github.com/gofiber/fiber/v2/middleware/logger"
//     "github.com/gofiber/fiber/v2/middleware/cache"
//     "github.com/gofiber/fiber/v2/middleware/basicauth"
// )

// func CORS() fiber.Handler {
//     return cors.New(cors.Config{AllowOrigins: "*"})
// }

// func LogAccess() fiber.Handler {
//     logApiAcccess := zlog.NewLogRollingFile(viper.GetString("app.logs.access")
//     return flog.New(flog.Config{
//         Format: "[${host}][${latency}][${method}][${status}] ${url}",
//         Output: logApiAcccess, //
//     })
// }

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

// func BasicAuth() fiber.Handler {
//     return basicauth.New(basicauth.Config{
//         Users: map[string]string{
//             "user":  "pass",
//         },
//     })
// }
