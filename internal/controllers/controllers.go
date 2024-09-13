package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"scoter-web-api/internal/models"
	"scoter-web-api/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

var UserCollection *mongo.Collection

func RegisterUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hash := sha256.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))

	user.ID = primitive.NewObjectID()

	_, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	var input models.User
	var storedUser models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hash := sha256.New()
	hash.Write([]byte(input.Password))
	input.Password = hex.EncodeToString(hash.Sum(nil))

	err := UserCollection.FindOne(context.TODO(), bson.M{"email": input.Email, "password": input.Password}).Decode(&storedUser)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Login successful", "user": storedUser})
}

func RentScooter(scooterID string, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(scooterID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID, "is_active": true}
	update := bson.M{"$set": bson.M{"rented_by": userID, "is_active": false}}

	result, err := repositories.ScooterCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error renting scooter:", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("No active scooter found to rent")
	}

	return nil
}
