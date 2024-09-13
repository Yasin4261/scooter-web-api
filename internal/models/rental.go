package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rental struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	ScooterID   primitive.ObjectID `bson:"scooter_id"`
	RentStart   time.Time          `bson:"rent_start"`
	RentEnd     time.Time          `bson:"rent_end,omitempty"`
	IsCompleted bool               `bson:"is_completed"`
}
