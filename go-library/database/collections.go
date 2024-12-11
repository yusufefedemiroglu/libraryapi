package database

import "go.mongodb.org/mongo-driver/mongo"

var (
	UserCollection *mongo.Collection
	BookCollection *mongo.Collection
)

// InitCollections initializes the MongoDB collections
func InitCollections(client *mongo.Client) {
	UserCollection = client.Database("library").Collection("users")
	BookCollection = client.Database("library").Collection("books")
}
