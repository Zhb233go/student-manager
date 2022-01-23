package controllers

import (
	. "StudentManager/middleware"
	"StudentManager/models"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Add(c *fiber.Ctx) error {
	var student = models.Student{}
	ctx := context.Background()
	s := c.Body()
	err := json.Unmarshal(s, &student)
	student.InsertOne(ctx, student)
	err = c.SendString("success add a single student!")
	if err != nil {
		return err
	}
	return nil
}

func QueryOne(c *fiber.Ctx) error {
	var student = models.Student{}
	ctx := context.Background()
	student.QueryOne(ctx, student)
	err := c.Status(200).JSON(student)
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info("query a single student!")
	return nil
}

func QueryAll(c *fiber.Ctx) error {
	var student = models.Student{}
	filter := bson.D{{"student", bson.D{{"$lte", 500}}}}
	ctx := context.Background()
	ss := student.QueryAll(ctx, filter)
	err := c.Status(200).JSON(ss)
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info("query all students")
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
