package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

func fiberServer(lc fx.Lifecycle) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(favicon.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success", "message": "go backend ok!"})
	})

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port 5000")
			go app.Listen(":5000")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app

}

func main() {
	fx.New(
		fx.Provide(),
		fx.Invoke(fiberServer),
	).Run()
}
