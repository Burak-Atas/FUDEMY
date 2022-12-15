package controller

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"new_CodingTime/token"
	"new_CodingTime/user/models"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddEduCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, err := r.Cookie("token")
	if err != nil {
		log.Println("hatalı token; Lütfen tekrar giriş yapınız")
		return
	}

	tok, _ := token.ValidateToken(cookie.Value)
	if tok.Uid == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uid := tok.Uid
	category := r.FormValue("category")
	title := r.FormValue("title")
	details := r.FormValue("details")
	price := r.FormValue("price")

	err1 := AddDatabase(ctx, uid, title, category, details, price)
	if !err1 {
		log.Fatal("eğitiminiz oluşturulamadı")
		return
	}

	log.Print("succesfull")

}

func AddDatabase(ctx context.Context, uid, title, category, details, price string) bool {

	var prod models.Product

	prod.Category = category
	prod.Details = details
	prod.Price = price
	prod.TeacherID = uid
	prod.Title = title

	created, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	prod.CreatedAt = created

	prod.ID = primitive.NewObjectID()

	_, err := prodCollection.InsertOne(ctx, prod)

	if err != nil {
		return false
	}

	return true
}

func DeleteEdu(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func AddCourse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	temp, err := template.ParseFiles("user/views/education/eduadd.html")
	if err != nil {

	}
	temp.Execute(w, nil)
}
