package controller

import (
	"Designweb/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.User(database.Client, "user")
var prodCollection *mongo.Collection = database.Product(*database.Client, "product")
var sepetCollection *mongo.Collection = database.Order(*database.Client, "sepet")
var cardCoollection *mongo.Collection = database.Card(*database.Client, "card")
var videoCollection *mongo.Collection = database.Videos(*database.Client, "video")
var commentCollection *mongo.Collection = database.Comment(*database.Client, "comment")
var buyercollectiong *mongo.Collection = database.Buyed(*database.Client, "buyed")
