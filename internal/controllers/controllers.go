package controllers

import (
	"log"
	"scoter-web-api/internal/models"
	"scoter-web-api/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllScooters(c *fiber.Ctx) error {
	scooters, err := repositories.GetAllScooters()
	if err != nil {
		log.Println("Error fetching scooters:", err) // Hata mesajını loglayın
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch scooters"})
	}
	return c.JSON(scooters)
}

func CreateScooter(c *fiber.Ctx) error {
	var scooter models.Scooter
	if err := c.BodyParser(&scooter); err != nil {
		log.Println("Error parsing input:", err) // Hata mesajını loglayın
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := repositories.CreateScooter(&scooter); err != nil {
		log.Println("Error creating scooter:", err) // Hata mesajını loglayın
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create scooter"})
	}

	return c.SendStatus(fiber.StatusCreated)
}
