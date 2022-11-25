// Package route api service
package route

import (
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/controller"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/controller/privatectl"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/controller/publicctl"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// Init api router
func Init(app fiber.Router) {
	// initailize controller
	controller.Init()

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
	publicctl.EndpointPing(rt)
}

func routePrivate(rt fiber.Router) {
	// app private route handler
	privatectl.EndpointPing(rt)

}
