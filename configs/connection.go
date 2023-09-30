package configs

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Client = ConnectDatabase()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	databaseName := os.Getenv("DatabaseName")
	return client.Database(databaseName).Collection(collectionName)
}
