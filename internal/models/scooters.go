package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Scooter struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
	IsActive  bool               `json:"is_active"`
}
