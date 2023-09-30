package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product MongoDB'deki "products" koleksiyonunu temsil eden yapıdır.
type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}
