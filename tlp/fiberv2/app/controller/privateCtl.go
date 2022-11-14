package controller

import "github.com/gofiber/fiber/v2"

// EndpointPing ping endpoint
//
//	@param r fiber.Router
func EndpointPrivatePing(r fiber.Router) {
	// init endpoint
	// ...
	// ...

	r.Get("/ping", getPing)
	// app more routers
}
