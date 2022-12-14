package config

import (
	"fmt"
	"runtime"

	"github.com/attapon-th/go-pkgs/zlog/log"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/model"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// NewFiberConfig fiber server initailize
func NewFiberConfig() fiber.Config {
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
	return config
}

// ListenString start server api
func ListenString() string {
	l := fmt.Sprintf("%s:%s", viper.GetString("app.listen"), viper.GetString("app.port"))
	log.Info().Str("Listener", l).Msg("API server started")
	return l
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
	resError := model.BaseResponse{}
	resError.Set(false, msgError)
	return ctx.Status(code).JSON(resError)
}
