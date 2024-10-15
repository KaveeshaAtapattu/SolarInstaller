package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProjectModel represents the structure in MongoDB
type ProjectModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ProjectName string             `bson:"projectName"`
	Location    string             `bson:"location"`
	DueDate     time.Time          `bson:"dueDate"`
	Status      bool               `bson:"status"`
}

// Project represents the API structure (client-facing)
type Project struct {
	ID         string    `json:"id"`
	ProjectName string    `json:"projectName"`
	Location    string    `json:"location"`
	DueDate     time.Time `json:"dueDate"`
	Status      bool      `json:"status"`
}
