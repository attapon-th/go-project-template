// Package publicctl publice endpoint
package publicctl

import (
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/model"
	"github.com/gofiber/fiber/v2"
)

// EndpointPing ping endpoint
//
//	@param r fiber.Router
func EndpointPing(r fiber.Router) {
	// init endpoint
	// ...
	// ...

	r.Get("/ping", getPing)
	// app more routers
}

func getPing(c *fiber.Ctx) error {
	res := model.BaseResponse{}
	res.Set(true, "Ping public endpoint successfully.")
	return c.JSON(res)
}
