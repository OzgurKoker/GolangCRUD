package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {
	//	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatalln(err)
	}
	context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golangAPI").Collection(collectionName)
}
