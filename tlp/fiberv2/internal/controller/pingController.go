package controller

import (
	"github.com/gofiber/fiber/v2"
)

func EndpointPing(rt fiber.Router, prefixUrl string) {
	r := rt.Group(prefixUrl, middleware.CORS(), middleware.LogAccess())

	r.Get("/ping", getPing)
}

func getPing(c *fiber.Ctx) error {
	return fiber.Mab{ok: true, msg: "ping succcess"}
}
