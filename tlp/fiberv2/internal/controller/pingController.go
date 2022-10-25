package controller

import (
	"github.com/gofiber/fiber/v2"
)

func EndpointPing(r fiber.Router) {
	r.Get("/ping", getPing)

}

func getPing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"ok": true, "msg": "ping succcess"})
}
