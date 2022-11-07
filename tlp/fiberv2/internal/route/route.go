// Package route api service
package route

import (
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// New api router
func New(r fiber.Router) {
	pathPrefix := viper.GetString("app.prefix")

	// Setup Middleware
	// r := app.Use(pathPrefix, middleware.CORS(), middleware.LogAccess())

	// Create Group route
	routePublic(r.Group(pathPrefix + "/public"))

	routePrivate(r.Group(pathPrefix))
}

func routePublic(rt fiber.Router) {
	//  app public route handler
	// import controller public
	controller.EndpointPing(rt)
}

func routePrivate(rt fiber.Router) {
	// app private route handler
	// import controller private

}
