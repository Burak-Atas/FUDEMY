package controller

import (
	"Designweb/token"
	"Designweb/user/models"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

type sendcard struct {
	Name string
	Cart []models.Card
}

func Mycard(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal("cookie verilerinde sorun var")
	}
	claims, _ := token.ValidateToken(cookie.Value)
	name := claims.First_name + " " + claims.Last_name

	var send sendcard
	send.Name = name

	var card []models.Card

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, _ := cardCoollection.Find(ctx, bson.M{"userid": claims.Uid})
	cursor.All(ctx, &card)
	send.Cart = card

	temp, _ := template.ParseFiles("user/views/card.html")
	temp.Execute(w, send)
}

func AddCard(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal("cookie verilerinde sorun var")
	}
	claims, _ := token.ValidateToken(cookie.Value)
	name := claims.First_name + " " + claims.Last_name

	var send sendcard
	send.Name = name

	number := r.PostFormValue("cardnumber")
	csv := r.PostFormValue("cvc")

	if ok := addcard_database(name, number, csv, claims.Uid); ok {
		log.Print("veri eklenirken hata oluştu")
		return
	}
	log.Println("kart verisi başarı ile eklendi")
	http.Redirect(w, r, "/mycard", http.StatusSeeOther)
}

func DeleteCard(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cardnumber := r.PostFormValue("cardnumber")
	fmt.Println(cardnumber)
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := cardCoollection.DeleteOne(ctx, bson.M{"cardnumber": cardnumber})
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("veriler başarı ile silindi")
	http.Redirect(w, r, "/mycard", http.StatusSeeOther)
}

func addcard_database(name, number, csv string, userid string) bool {

	var cardmodel models.Card
	cardmodel.Name = name
	cardmodel.Cardnumber = number
	cardmodel.Csv = csv
	cardmodel.UserId = userid

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, _ := cardCoollection.CountDocuments(ctx, bson.M{"cardnumber": cardmodel.Cardnumber})

	if count > 0 {
		log.Println("bu kart zaten mevcut")
		return true
	}

	_, err := cardCoollection.InsertOne(ctx, cardmodel)
	return err != nil
}
