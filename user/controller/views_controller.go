package controller

import (
	"context"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"new_CodingTime/database"
	"new_CodingTime/token"
	"new_CodingTime/user/models"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"tawesoft.co.uk/go/dialog"
)

//şifre hashleme
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

//Pasaport eşleme kontrolü
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//databases Erişim
var UserCollection *mongo.Collection = database.User(database.Client, "user")
var prodCollection *mongo.Collection = database.Product(*database.Client, "product")

//home page
func Index_html(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")
	temp, _ := template.ParseFiles("user/views/index.html")

	if err != nil {
		IsAdmin := false
		temp.Execute(w, IsAdmin)
		return
	}

	claims, _ := token.ValidateToken(cookie.Value)
	temp.Execute(w, claims)

}

//	Login page
func Login_html(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	temp, _ := template.ParseFiles("user/views/login.html")

	temp.Execute(w, nil)
}

func Blog_html(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var blogList []models.Blog
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := prodCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("veriler alınırken hata oluştu", err)
		return
	}
	if err = cursor.All(ctx, &blogList); err != nil {
		log.Fatal(err)
	}
	paths := []string{
		"/user/views/blog/blog.html",
	}
	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))

	t.Execute(w, blogList)
}

func DatabesesIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var product []models.Product
	_, err := r.Cookie("token")
	if err != nil {
		log.Println("token hatalı lütfen tekrar giriş yapınız")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	//token.ValidateToken(cookie.Value)

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := prodCollection.Find(ctx, bson.M{"category": "database"})
	if err != nil {
		log.Panicln("veritabanında veriler alınırken hata oluştu")
		return
	}

	temp := template.Must(template.ParseFiles("user/views/education/databases.html"))

	cursor.All(ctx, &product)
	temp.Execute(w, product)

}

func UsLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	var user models.User
	var foundUser models.User

	user.Email = email
	user.Password = HashPassword(password)

	var filter = bson.M{"email": email}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := UserCollection.FindOne(ctx, filter).Decode(&foundUser)
	if err != nil {

	}

	if foundUser.Password != password {

		dialog.Alert("şifre hatalı")

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return

	}

	cookie := &http.Cookie{
		Name:   "token",
		Value:  foundUser.Token,
		MaxAge: 30000,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func UsSignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var user models.User

	user.Email = r.FormValue("email")
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Password = r.FormValue("password")

	user.Type = "STUDENT"
	user.ID = primitive.NewObjectID()
	user.UserID = ""

	user.CreatedAt, _ = (time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)))
	user.UpdatedAt, _ = (time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)))
	Token, _ := token.CreateToken(user.Email, user.FirstName, user.LastName, user.UserID)

	user.Token = Token

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, _ := UserCollection.CountDocuments(ctx, bson.M{
		"email": user.Email,
	})

	if count > 1 {

		dialog.Alert("Veri tabanına zaten kayıtlı!")

		http.Redirect(w, r, "/login", http.StatusSeeOther)

		return
	}

	result, err := UserCollection.InsertOne(ctx, user)
	log.Print(result)
	if err != nil {
		log.Println("kullanıcı verileri eklenirken hata oluştu")
		return
	}

	log.Print("redirect")

	cookie := &http.Cookie{
		Name:   "token",
		Value:  user.Token,
		MaxAge: 30000,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func BackEndIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	temp, err := template.ParseFiles("")
	if err != nil {

	}
	var product []models.Product
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{"category": "backend"}
	cursor, _ := prodCollection.Find(ctx, filter)
	defer cancel()

	cursor.All(ctx, &product)
	temp.Execute(w, product)
}

func EtichalIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

//öğretmene ait kurslar
func TeacherAdmin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, err := r.Cookie("token")
	if err != nil {

	}
	claims, _ := token.ValidateToken(cookie.Value)
	if claims.Uid == "" {

	}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"TeacherID": claims.Uid}
	cursor, err := prodCollection.Find(ctx, filter)
	if err != nil {

	}

	var prod []models.Product
	cursorErr := cursor.All(ctx, &prod)
	if cursorErr != nil {

	}

	temp := template.Must(template.ParseFiles(""))
	temp.Execute(w, nil)
}

var Sepetcollection = *database.Order(*database.Client, "sepet")

func AddCard(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("token")
	if err != nil {

	}

	claims, _ := token.ValidateToken(cookie.Value)
	if claims.Uid == "" {

	}
	id := r.FormValue("id")
	var model models.Sepet

	model.UserId = claims.Uid
	model.EduId = id
	model.ID = primitive.NewObjectID()
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Sepetcollection.InsertOne(ctx, model)
	fmt.Println("sepete başarı ile eklendi")
	http.Redirect(w, r, "/sepet", http.StatusSeeOther)
}
