// Package route api service
package route

import (
	"github.com/attapon-th/go-project-template/tpl/fiberv2/app/controller"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// New api router
func New(app fiber.Router) {
	// initailize controller
	controller.New()

	pathPrefix := viper.GetString("app.prefix")

	// Setup Middleware all route
	r := app.Use(pathPrefix, middleware.CORS(), middleware.LogAccess())

	// Create Group route Public
	routePublic(r.Group(pathPrefix + "/public"))

	// Create Group route Private
	routePrivate(r.Group(pathPrefix+"/private", middleware.BasicAuth()))
}

func routePublic(rt fiber.Router) {
	//  app public route handler
	controller.EndpointPing(rt)
}

func routePrivate(rt fiber.Router) {
	// app private route handler
	controller.EndpointPrivatePing(rt)

}
