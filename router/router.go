package router

import (
	"fiber-demo/logs"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupRouter() *fiber.App {
	app := fiber.New()
	logs.Logger(app)
	// 定义全局的中间件
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First 2222handler")
		return c.Next()
	})

	//对参数的校验
	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("what you want to do?")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Println("listen error: ", err)
	}
	return app
}
