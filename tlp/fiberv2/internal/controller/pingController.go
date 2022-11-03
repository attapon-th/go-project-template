package controller

import (
	"github.com/gofiber/fiber/v2"
)

// PingController is a controller for ping
type PingController struct {
	OK  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func EndpointPing(r fiber.Router) {
	r.Get("/ping", getPing)

}

func getPing(c *fiber.Ctx) error {
	return c.JSON(PingController{OK: true, Msg: "ping succcess"})
}
