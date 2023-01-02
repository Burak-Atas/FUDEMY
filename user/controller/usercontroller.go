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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// şifre hashleme
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// Pasaport eşleme kontrolü
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("login --> %s\n", r.Method)
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("user/views/login.html")

		temp.Execute(w, nil)
	} else {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		var user models.User
		var foundUser models.User

		user.Email = email
		user.Password = HashPassword(password)

		var filter = bson.M{"email": email}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := userCollection.FindOne(ctx, filter).Decode(&foundUser)
		if err != nil {
			log.Fatal(err)
		}

		if foundUser.Password != password {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		cookie := &http.Cookie{
			Name:   "token",
			Value:  foundUser.Token,
			MaxAge: 3000,
		}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}

func SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var user models.User
	r.ParseForm()
	user.Email = r.FormValue("email")
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Password = r.FormValue("password")

	user.Type = "STUDENT"
	user.ID = primitive.NewObjectID()
	user.UserID = primitive.NewObjectID().Hex()

	user.CreatedAt, _ = (time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)))
	user.UpdatedAt, _ = (time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)))
	Token, _ := token.CreateToken(user.Email, user.FirstName, user.LastName, user.UserID)

	user.Token = Token

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, _ := userCollection.CountDocuments(ctx, bson.M{
		"email": user.Email,
	})

	if count > 1 {

		http.Redirect(w, r, "/login", http.StatusSeeOther)

		return
	}

	result, err := userCollection.InsertOne(ctx, user)
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
