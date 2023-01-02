package controller

import (
	"Designweb/token"
	"Designweb/user/models"
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func UploadVideo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("method:", r.Method)

	videoisim := r.FormValue("title")

	// Dosya türünü ve boyutunu kontrol et
	r.ParseMultipartForm(120 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Dosya yüklenirken bir hata oluştu", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if handler.Size > (150 << 20) {
		http.Error(w, "Dosya boyutu çok büyük", http.StatusRequestEntityTooLarge)
		return
	}
	videourl := "user/asset/video" + "/" + videoisim + ".mp4"
	// Dosya adını tekrar kullanılmamasını sağla
	f, err := os.OpenFile(videourl, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Dosya yüklenirken bir hata oluştu", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	// Yükleme işlemini tamamla

	AddDatabasevideo(videourl, videoisim)

	http.Redirect(w, r, "/addcourse", http.StatusSeeOther)
}

func AddDatabasevideo(url, title string) bool {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var videoModel models.Video
	videoModel.VideoTitle = title
	videoModel.VideoUrl = url
	videoModel.EduId = edueid

	videoModel.ID = primitive.NewObjectID()
	videoModel.VideoId = primitive.NewObjectID().Hex()

	_, err := videoCollection.InsertOne(ctx, videoModel)
	return err != nil
}

type SendVideo struct {
	Name     string
	Video    models.Video
	VideoUrl string
}

func Eduvideo(w http.ResponseWriter, r *http.Request, router httprouter.Params) {
	cookie, err := r.Cookie("token")

	var send SendVideo

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	query := r.URL.Query()
	eduid := query["id"][0]
	claims, _ := token.ValidateToken(cookie.Value)
	name := claims.First_name + claims.Last_name
	send.Name = name

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var v models.Video

	videoCollection.FindOne(ctx, bson.M{"videoid": eduid}).Decode(&v)

	send.Video = v
	send.VideoUrl = "/" + send.Video.VideoUrl

	fmt.Println(send.VideoUrl)
	temp, _ := template.ParseFiles("user/views/video.html")
	temp.Execute(w, send)
}
