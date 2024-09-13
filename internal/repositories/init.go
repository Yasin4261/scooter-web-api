package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ScooterCollection *mongo.Collection
	UserCollection    *mongo.Collection
	RentalCollection  *mongo.Collection
	// Diğer koleksiyonlar burada tanımlanabilir
)

// InitRepositories initializes all MongoDB collections
func InitRepositories(db *mongo.Database) {
	ScooterCollection = db.Collection("scooters")
	UserCollection = db.Collection("users")
	RentalCollection = db.Collection("rentals")
	// Diğer koleksiyonlar burada başlatılabilir
}
