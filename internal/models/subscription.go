package models

import "time"

type Subscription struct {
	Id          string    `bson:"_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Price       string    `bson:"price"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	Lessons     []Lesson  `bson:"lessons"`
}
