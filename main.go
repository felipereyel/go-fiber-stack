package main

import (
	"github.com/felipereyel/PROJECT_NAME/pkgs/config"
	"github.com/felipereyel/PROJECT_NAME/pkgs/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Load()
	app := fiber.New()

	app.Use(cors.New())
	app.Use("/healthcheck", handler.HealthCheck)
	app.Use("/protected/:hookId", handler.ScopeAuth(handler.HookScope), handler.HealthCheck)

	err := app.Listen(":" + config.Config("PORT"))
	if err != nil {
		panic(err.Error())
	}
}
