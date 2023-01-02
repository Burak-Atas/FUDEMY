package controller

import (
	"Designweb/token"
	"Designweb/user/models"
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

type Senadata struct {
	Name    string
	Product []models.Product
	Video   []models.Video
}
type edu_dvideo struct {
	Name         string
	Product_name string
	Details      string
	Eduid        string
	Video        []models.Video
	Yorumlar     []models.Yorumlar
}

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, err := r.Cookie("token")
	temp, _ := template.ParseFiles("user/views/index.html")
	if err != nil {
		temp.Execute(w, nil)
		return
	}
	claims, _ := token.ValidateToken(cookie.Value)
	name := claims.First_name + " " + claims.Last_name
	var product []models.Product
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, _ := prodCollection.Find(ctx, bson.M{})
	cursor.All(ctx, &product)

	var data Senadata

	data.Name = name
	data.Product = product

	temp.Execute(w, data)

}

func Education(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var cookie, err = r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	claims, _ := token.ValidateToken(cookie.Value)
	name := claims.First_name + " " + claims.Last_name

	var sendata Senadata

	id := r.FormValue("id")
	var product models.Product
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"ProductId": id}

	prodCollection.FindOne(ctx, filter).Decode(product)

	temp, _ := template.ParseFiles("user/views/ed.html")

	sendata.Name = name
	sendata.Product[0] = product

	temp.Execute(w, sendata)
}
