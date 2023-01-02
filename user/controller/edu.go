package controller

import (
	"Designweb/token"
	"Designweb/user/models"
	"context"
	"errors"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Sendd struct {
	Name    string
	Product models.Product
}

var (
	edueid, eduisim string
)

func AddEducation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")

	var send Sendd
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	claims, _ := token.ValidateToken(cookie.Value)

	if r.Method == "GET" {
		name := claims.First_name + " " + claims.Last_name
		send.Name = name
		send.Product.ProductId = edueid

		temp, _ := template.ParseFiles("user/views/addcours.html")
		temp.Execute(w, send)

	} else {

		title := r.PostFormValue("title")
		details := r.PostFormValue("details")
		category := r.PostFormValue("category")
		price := r.PostFormValue("price")
		userid := claims.Uid
		url := ""
		id := primitive.NewObjectID()
		ProductId := primitive.NewObjectID().Hex()

		edueid = ProductId
		eduisim = title

		ok := add_edu_databases(title, details, category, price, url, userid, id, ProductId)
		if ok {
			log.Print("hata oluştu")
			return
		}

		log.Println("Videonuz başarı ile eklendi")
		http.Redirect(w, r, "/addcourse", http.StatusSeeOther)
	}
}

func add_edu_databases(title, details, category, price, url string, userid string, id primitive.ObjectID, productid string) bool {
	var product models.Product
	product.Title = title
	product.Details = details
	product.Category = category
	product.Price = price
	product.Url = url
	product.ProductId = productid
	product.ID = id

	date, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	product.CreatedAt = date
	product.TeacherID = userid

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := prodCollection.InsertOne(ctx, product)
	return err != nil

}

func MyCourse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")

	var send Senadata
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	claims, _ := token.ValidateToken(cookie.Value)
	name := claims.First_name + claims.Last_name
	send.Name = name

	var product []models.Product
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := prodCollection.Find(ctx, bson.M{"teacherid": claims.Uid})
	defer cancel()

	if err != nil {
		log.Fatal(errors.New("veri tabanından kurslarınız çekilirken hata oluştu"))
		return
	}

	cursor.All(ctx, &product)
	send.Product = product

	temp, _ := template.ParseFiles("user/views/mycourse.html")
	temp.Execute(w, send)

}

func Course(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")

	var send edu_dvideo
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	claims, _ := token.ValidateToken(cookie.Value)
	name := claims.First_name + claims.Last_name
	send.Name = name

	var video []models.Video
	var product models.Product
	id := r.PostFormValue("id")
	log.Println(id)

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	prodCollection.FindOne(ctx, bson.M{"productid": id}).Decode(&product)
	c, _ := commentCollection.Find(ctx, bson.M{"EduId": id})

	c.All(ctx, &send.Yorumlar)
	cursor, _ := videoCollection.Find(ctx, bson.M{"eduid": id})
	cursor.All(ctx, &video)
	send.Video = video
	send.Product_name = product.Title
	send.Details = product.Details
	send.Eduid = id
	temp, _ := template.ParseFiles("user/views/ed.html")
	temp.Execute(w, send)
}

func EduSil(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.PostFormValue("id")
	log.Println("id", id)
	prodCollection.DeleteOne(ctx, bson.M{"productid": id})
	videoCollection.DeleteMany(ctx, bson.M{"eduid": id})

	http.Redirect(w, r, "/mycourse", http.StatusSeeOther)

}

func EduGuncelle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, err := r.Cookie("token")

	var send Sendd
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	claims, _ := token.ValidateToken(cookie.Value)
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	id := r.PostFormValue("eduid")
	if r.Method == "GET" {
		name := claims.First_name + " " + claims.Last_name
		send.Name = name
		prodCollection.FindOne(ctx, bson.M{"productid": id})
		defer cancel()
		temp, _ := template.ParseFiles("user/views/eduguncelle.html")
		temp.Execute(w, send)

	} else {

		title := r.PostFormValue("title")
		details := r.PostFormValue("details")
		price := r.PostFormValue("price")

		var prod models.Product
		prod.Title = title
		prod.Details = details
		prod.Price = price

		defer cancel()
		prodCollection.UpdateByID(ctx, id, prod)
		http.Redirect(w, r, "/mycourse", http.StatusSeeOther)

	}
}
