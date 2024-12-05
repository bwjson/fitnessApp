package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	UserID        primitive.ObjectID `json:"userID" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty" binding:"required"`
	Email         string             `json:"email" bson:"email,omitempty"`
	Password      string             `json:"password" bson:"password,omitempty"`
	RegisteredAt  time.Time          `json:"registeredAt" bson:"registeredAt,omitempty"`
	LastVisitAt   time.Time          `json:"lastVisitAt" bson:"lastVisitAt,omitempty"`
	Subscriptions []Subscription     `json:"subscriptions" bson:"subscriptions,omitempty"`
}
