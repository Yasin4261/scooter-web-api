package controllers

import (
	"context"
	"net/http"
	"time"

	"scoter-web-api/internal/models"
	"scoter-web-api/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartRental(c *fiber.Ctx) error {
	type Request struct {
		ScooterID string `json:"scooter_id"`
		UserID    string `json:"user_id"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// ScooterID ve UserID string türünden primitive.ObjectID'ye dönüştürülüyor
	scooterID, err := primitive.ObjectIDFromHex(request.ScooterID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid scooter ID"})
	}

	userID, err := primitive.ObjectIDFromHex(request.UserID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	rental := models.Rental{
		ID:          primitive.NewObjectID(),
		ScooterID:   scooterID,
		UserID:      userID,
		RentStart:   time.Now(),
		IsCompleted: false,
	}

	err = repositories.CreateRental(&rental)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to start rental"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Rental started successfully"})
}

func GetUserRentals(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	rentals, err := repositories.GetUserRentals(userID.Hex())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch rentals"})
	}

	return c.JSON(rentals)
}

func CompleteRental(c *fiber.Ctx) error {
	rentalID := c.Params("id")

	objectRentalID, err := primitive.ObjectIDFromHex(rentalID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid rental ID"})
	}

	filter := bson.M{"_id": objectRentalID}
	update := bson.M{"$set": bson.M{
		"rent_end":     time.Now(),
		"is_completed": true,
	}}

	_, err = repositories.RentalCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to complete rental"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Rental completed successfully"})
}
