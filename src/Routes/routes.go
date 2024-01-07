package routes

import (
	"github.com/gofiber/fiber/v2"
	models "root/src/Model"
	database "root/src/Database"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("епта работает")
}

func AddUser(c *fiber.Ctx) error {
	dto := new(models.Dto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&dto)

	return c.Status(200).JSON(dto)
}

func AllUsers(c *fiber.Ctx) error {
	data := []models.Dto{}
	database.DB.Db.Find(&data)

	return c.Status(200).JSON(data)
}