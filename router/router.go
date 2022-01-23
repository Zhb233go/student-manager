package router

import (
	"StudentManager/controllers"
	. "StudentManager/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

// SetupRouter 路由初始化
func SetupRouter() *fiber.App {
	// Initialize standard Go html template engine
	engine := html.New("./template/user", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		//ViewsLayout: "layout/main",
	})
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(recover.New())
	//对参数的校验
	//app.Get("/:name?", controllers.Hello)
	manager := app.Group("/manager")
	{
		manager.Post("/student", controllers.Add)
		manager.Get("/student", controllers.QueryOne)
		manager.Get("/all_student", controllers.QueryAll)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	if err := app.Listen(":3000"); err != nil {
		Ln.Fatal(err)
	}
	Ln.Info("监听在:3000 端口")
	return app
}
