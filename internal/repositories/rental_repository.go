package repositories

import (
	"context"
	"log"
	"time"

	"scoter-web-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Scooter kiralama işlemi kaydetme
func CreateRental(rental *models.Rental) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := RentalCollection.InsertOne(ctx, rental)
	if err != nil {
		log.Printf("Error creating rental: %v\n", err)
		return err
	}
	return nil
}

// Kullanıcının aktif kiralamalarını getirme
func GetUserRentals(userID string) ([]models.Rental, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	cursor, err := RentalCollection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error fetching rentals: %v\n", err)
		return nil, err
	}

	var rentals []models.Rental
	if err := cursor.All(ctx, &rentals); err != nil {
		log.Printf("Error decoding rentals: %v\n", err)
		return nil, err
	}

	return rentals, nil
}
