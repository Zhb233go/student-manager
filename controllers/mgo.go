package controllers

import (
	. "StudentManager/middleware"
	"StudentManager/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func Save(c *fiber.Ctx) error {
	var student = models.Student{}
	s := c.Body()
	err := json.Unmarshal(s, &student)
	err = student.Insert()
	if err != nil {
		Ln.Error(err)
	}
	c.Status(200).SendString("success add a single student!student.Uid:" + string(student.Uid))
	return nil
}

func FindAll(c *fiber.Ctx) error {
	var student = models.Student{}
	all, err := student.All()
	if err != nil {
		Ln.Error(err)
	}
	c.Status(200).JSON(all)
	return nil
}

func Hello(c *fiber.Ctx) error {
	if c.Params("name") != "" {
		Ln.Info("hello world")
		return c.SendString("Hello " + c.Params("name"))
		// => Hello john
	}
	Ln.Info("hello,who are you?")
	return c.SendString("what you want to do?")
}
