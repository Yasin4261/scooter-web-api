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

func UpdateScooterLocation(c *fiber.Ctx) error {
	type Request struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	id := c.Params("id")

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := repositories.UpdateScooterLocation(id, request.Latitude, request.Longitude)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update scooter location"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func UpdateScooterStatus(c *fiber.Ctx) error {
	type Request struct {
		IsActive bool `json:"is_active"`
	}

	id := c.Params("id")

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := repositories.UpdateScooterStatus(id, request.IsActive)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update scooter status"})
	}

	return c.SendStatus(fiber.StatusOK)
}
