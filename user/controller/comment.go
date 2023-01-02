package controller

import (
	"Designweb/token"
	"Designweb/user/models"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Addcomment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	claims, _ := token.ValidateToken(cookie.Value)

	comment := r.PostFormValue("yorum")
	eduid := r.PostFormValue("eduid")
	userid := claims.Uid

	ok := Addcomment_Database(comment, eduid, userid)
	if ok {
		log.Println("yorum gönderilirken hata oluştu")
		return
	}
	log.Println("yorum ekleme işlemi başarılı")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Addcomment_Database(comment, eduid, userid string) bool {
	var commentModel models.Yorumlar
	commentModel.CommentDetails = comment
	commentModel.EduId = eduid
	commentModel.UserId = userid

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := commentCollection.InsertOne(ctx, commentModel)
	return err != nil
}
