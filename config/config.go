package config

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

const mongoURI = "mongodb+srv://kaveevenuranga:fQAVkkt2aRdN57LG@solarinstaller.jx4vz.mongodb.net/?retryWrites=true&w=majority"

func InitDB() {
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	DB = client.Database("SolarInstaller")
	log.Println("Connected to MongoDB!")
}
