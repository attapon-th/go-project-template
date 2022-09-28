package internal

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

var (
	// Version of builds
	Version string

	// Build hashing
	Build string

	// Timestamp of build
	Timestamp string
)

func Serv() {
	app := NewFiber()
	// controller.EndpointPing(app, viper.GetString("app.prefix"))
}

func NewFiber() *fiber.App {
	app := fiber.New
	config := fiber.Config{
		AppName:       "fiber-api",
		Prefork:       false,
		StrictRouting: true,
		CaseSensitive: true,
		Immutable:     false,
		BodyLimit:     4 * 1024 * 1024, // 4mb body size
		Concurrency:   256 * 1024,
		JSONEncoder:   jsoniter.Marshal,
		JSONDecoder:   jsoniter.Unmarshal,
		ErrorHandler:  ErrorHandlerResponseJSON,
	}
	if nc := viper.GetInt("app.cpu"); nc == 1 || nc < 0 {
		config.Prefork = false
	} else if nc == 0 {
		config.Prefork = true
	} else if nc > 1 {
		config.Prefork = true
		runtime.GOMAXPROCS(nc)
	}
	app := fiber.New(config)
	return app
}

func ErrorHandlerResponseJSON(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msgError := "Error!!!"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	if err != nil {
		msgError = err.Error()
	}
	resError := fiber.Map{
		"msg": msgError,
		"ok":  false,
	}
	return ctx.Status(code).JSON(resError)
}
