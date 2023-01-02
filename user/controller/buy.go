package controller

import (
	"Designweb/token"
	"Designweb/user/models"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func Buy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	claims, _ := token.ValidateToken(cookie.Value)

	id := r.PostFormValue("id")
	var buy models.Control

	buy.ID = primitive.NewObjectID()
	buy.ProductId = id
	buy.UserID = claims.Uid
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, _ := cardCoollection.CountDocuments(ctx, bson.M{"userid": claims.Uid})
	if c < 1 {
		log.Println("lütfen bir kart ekleyin")
		http.Redirect(w, r, "/mycard", http.StatusSeeOther)
	}
	buyercollectiong.InsertOne(ctx, buy)
	log.Println("satın alma başarılı")
	sepetCollection.DeleteOne(ctx, bson.M{"userid": claims.Uid})
	http.Redirect(w, r, "/order", http.StatusSeeOther)
}
