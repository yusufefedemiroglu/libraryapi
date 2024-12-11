package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the context is canceled to free resources

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err // Return error instead of terminating the program
	}

	// Test the connection by pinging the MongoDB server
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB!")
	DB = client 
	return client, nil
}

// Disconnect disconnects the MongoDB client
func Disconnect() {
	if DB != nil {
		err := DB.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("MongoDB bağlantısını kapatırken hata: %v", err)
		}
		log.Println("MongoDB bağlantısı kapatıldı.")
	}
}
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("library").Collection(collectionName)
}
