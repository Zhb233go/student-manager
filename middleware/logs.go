package logs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

// Logger 日志打印
func Logger(app *fiber.App) func() {
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Format:     "${time}  ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "local",
		Output:     file,
	}))
	//文件关闭接口返回上一层
	return func() {
		file.Close()
	}
}
