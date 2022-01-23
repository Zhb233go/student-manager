package router

import (
	"StudentManager/controllers"
	. "StudentManager/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRouter 路由初始化
func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(recover.New())
	//对参数的校验
	app.Get("/:name?", controllers.Hello)
	manager := app.Group("/manager")
	{
		manager.Post("/student", controllers.Add)
		//manager.Post("/student2", controllers.Save)
		manager.Get("/student", controllers.QueryOne)
		manager.Get("/allstudent", controllers.QueryAll)
		//manager.Get("/allstudent2", controllers.FindAll)
	}
	if err := app.Listen(":3000"); err != nil {
		Ln.Fatal(err)
	}
	Ln.Info("监听在:3000 端口")
	return app
}
