package controllers

import (
	. "StudentManager/middleware"
	"StudentManager/models"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var student = models.Student{}

func Add(c *fiber.Ctx) error {
	ctx := context.Background()
	s := c.Body()
	err := json.Unmarshal(s, &student)
	if err != nil {
		Ln.Error(err)
	}
	models.Insert(ctx, bson.M{"student": student})
	c.SendString("success add a single student!")
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
