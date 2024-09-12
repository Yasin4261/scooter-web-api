package repositories

import (
	"context"
	"log"
	"scoter-web-api/internal/config"
	"scoter-web-api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var scooterCollection *mongo.Collection

func InitScooterRepository() {
	scooterCollection = config.DB.Collection("scooters")
}

func GetAllScooters() ([]models.Scooter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var scooters []models.Scooter
	cursor, err := scooterCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error fetching scooters:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var scooter models.Scooter
		if err := cursor.Decode(&scooter); err != nil {
			log.Println("Error decoding scooter:", err)
			return nil, err
		}
		scooters = append(scooters, scooter)
	}

	return scooters, nil
}

func CreateScooter(scooter *models.Scooter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := scooterCollection.InsertOne(ctx, scooter)
	if err != nil {
		log.Println("Error creating scooter:", err)
		return err
	}

	return nil
}
