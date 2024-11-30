package models

import "time"

type Lesson struct {
	Id          string     `bson:"_id"`
	Categories  []Category `bson:"categories"`
	Title       string     `bson:"title"`
	Description string     `bson:"description"`
	CreatedAt   time.Time  `bson:"createdAt"`
	UpdatedAt   time.Time  `bson:"updatedAt"`
}
