package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// UserModel represents the user structure in MongoDB
type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

// UserLogin represents the login request payload
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
