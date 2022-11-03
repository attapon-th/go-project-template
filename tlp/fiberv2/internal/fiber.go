package internal

import (
	"fmt"
	"runtime"

	"github.com/attapon-th/go-pkgs/zlog/log"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (

	// AppFiber fiber app
	AppFiber *fiber.App

	// Version of builds
	Version string

	// Build hashing
	Build string

	// Timestamp of build
	Timestamp string

	// InitLoader func start loader manual add
	InitLoader []func()
)

// NewFiber fiber server initailize
func NewFiber() {
	for _, fn := range InitLoader {
		fn()
	}
	config := fiber.Config{
		AppName:       "fiber-api",
		Prefork:       false,
		StrictRouting: true,
		CaseSensitive: true,
		Immutable:     false,
		BodyLimit:     4 * 1024 * 1024, // 4mb body size
		Concurrency:   256 * 1024,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  errorHandlerResponseJSON,
	}

	if nc := viper.GetInt("app.cpu"); nc > 1 {
		config.Prefork = true
		runtime.GOMAXPROCS(nc)
	} else {
		config.Prefork = false
	}
	AppFiber = fiber.New(config)
}

// Serv start server api
// *** Uncomment // route.New(app) **
func Serv() {
	// route.New(app)
	l := fmt.Sprintf("%s:%s", viper.GetString("app.listen"), viper.GetString("app.port"))
	log.Info().Str("Listener", l).Msg("API server started")
	AppFiber.Listen(l)
}

func errorHandlerResponseJSON(ctx *fiber.Ctx, err error) error {
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
