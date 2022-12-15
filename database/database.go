package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBSet()

func DBSet() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	err2 := client.Connect(ctx)
	defer cancel()

	if err2 != nil {
		log.Println("users not connection the databases")
	}

	client.Ping(context.TODO(), nil)

	log.Println("veri tabanına bağlanıldı")
	return client
}

func User(client *mongo.Client, collname string) *mongo.Collection {
	return (*mongo.Collection)(client.Database("E-TİCARET").Collection(collname))
}

func Product(client mongo.Client, collname string) *mongo.Collection {
	return (*mongo.Collection)(client.Database("E-TİCARET").Collection(collname))
}

func Order(client mongo.Client, collname string) *mongo.Collection {
	return (*mongo.Collection)(client.Database("E-TİCARET").Collection(collname))
}
