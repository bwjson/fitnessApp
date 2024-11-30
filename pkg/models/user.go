package models

import "time"

type User struct {
	Id            string         `bson:"_id"`
	Name          string         `bson:"name"`
	Email         string         `bson:"email"`
	Password      string         `bson:"password"`
	RegisteredAt  time.Time      `bson:"registeredAt"`
	LastVisitAt   time.Time      `bson:"lastVisitAt"`
	Subscriptions []Subscription `bson:"subscriptions"`
}
