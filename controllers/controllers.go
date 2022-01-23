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
	if err != nil {
		Ln.Error(err)
	}
	student.InsertOne(ctx, bson.D{{"student", student}})
	c.SendString("success add a single student!")
	return nil
}

func QueryOne(c *fiber.Ctx) error {
	var student = models.Student{}
	ctx := context.Background()
	student.QueryOne(ctx, student)
	c.Status(200).JSON(student)
	Ln.Info("query a single student!")
	return nil
}

func QueryAll(c *fiber.Ctx) error {
	var student = models.Student{}
	ctx := context.Background()
	s := student.QueryAll(ctx, student)
	c.Status(200).JSON(s)
	Ln.Info("query all students")
	return nil
}
