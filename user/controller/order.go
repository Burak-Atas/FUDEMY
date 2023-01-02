package controller

import (
	"Designweb/token"
	"Designweb/user/models"
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func Addsepet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	claims, _ := token.ValidateToken(cookie.Value)
	id := r.PostFormValue("id")
	userid := claims.Uid
	var sepet models.Sepet

	sepet.ID = primitive.NewObjectID()
	sepet.EduId = id
	sepet.UserId = userid

	var ctx, cancel = context.WithTimeout(context.Background(), 19*time.Second)
	count, _ := sepetCollection.CountDocuments(ctx, bson.M{"eduid": id})
	if count > 0 {
		http.Redirect(w, r, "/order", http.StatusSeeOther)
	}

	sepetCollection.InsertOne(ctx, sepet)
	defer cancel()

	http.Redirect(w, r, "/order", http.StatusSeeOther)

	log.Print("sepete başarı ile eklendi")
}

type SendSepet struct {
	Name    string
	Product []models.Product
}

func Sepet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	claims, _ := token.ValidateToken(cookie.Value)

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, _ := sepetCollection.Find(ctx, bson.M{"userid": claims.Uid})
	var sepet []models.Sepet
	cursor.All(ctx, &sepet)
	var pro models.Product
	var product []models.Product
	for i := 0; i < len(sepet); i++ {
		prodCollection.FindOne(ctx, bson.M{"productid": sepet[i].EduId}).Decode(&pro)
		product = append(product, pro)
	}

	var m SendSepet
	m.Name = claims.First_name + " " + claims.Last_name
	m.Product = product

	temp, _ := template.ParseFiles("user/views/order.html")
	temp.Execute(w, m)
}
