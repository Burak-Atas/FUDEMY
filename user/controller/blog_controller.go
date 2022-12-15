package controller

import (
	"context"
	"log"
	"net/http"
	"new_CodingTime/database"
	"new_CodingTime/token"
	"new_CodingTime/user/models"

	"time"

	"github.com/julienschmidt/httprouter"
)

func Addblog(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, _ := r.Cookie("token")
	claims, _ := token.ValidateToken(cookie.Value)

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	title := r.FormValue("title")
	content := r.FormValue("content")

	msg := Add(ctx, title, content, claims.Uid)
	log.Println(msg)

}

func DeleteBlog(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

var blogCollectioon = database.Product(*database.Client, "blogs")

func Add(ctx context.Context, title, content, userId string) string {
	var bl models.Blog
	bl.Blog_Title = title

	date, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	bl.CreateDate = date
	blogCollectioon.InsertOne(ctx, bl)
	msg := "Blog verisi eklenmiştir : Blog" + bl.Blog_Title
	return msg
}

func Delete() {

}
