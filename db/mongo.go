package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kms/config"
	"log"
)

var MongoClient *mongo.Client
var MongoDb *mongo.Database

func ConnectDatabase() {
	dbConfig := config.DBConfig.Database

	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%d", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	clienOption := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.TODO(), clienOption)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	MongoClient = client
	MongoDb = client.Database(dbConfig.DatabaseName)
}

func DisConnectDatabase() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Failed to disconnect MongoDB: %v", err)
	}
	fmt.Println("Disconnected MongoDB!")
}
