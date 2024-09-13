package repositories

import (
	"context"
	"errors"
	"log"
	"scoter-web-api/internal/models"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllScooters() ([]models.Scooter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var scooters []models.Scooter
	cursor, err := ScooterCollection.Find(ctx, bson.M{})
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

	_, err := ScooterCollection.InsertOne(ctx, scooter)
	if err != nil {
		log.Println("Error creating scooter:", err)
		return err
	}

	return nil
}

func UpdateScooterLocation(id string, latitude float64, longitude float64) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"latitude":  latitude,
		"longitude": longitude,
	}}
	result, err := ScooterCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating scooter location:", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("No scooter found to update")
	}

	return nil
}

func UpdateScooterStatus(id string, isActive bool) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"is_active": isActive}}

	result, err := ScooterCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating scooter status: ", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("No scooter found to update")
	}

	return nil
}
