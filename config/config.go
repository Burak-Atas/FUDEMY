package config

import (
	"Designweb/user/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router_Config() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", controller.Index)

	//kullanıcı giriş kayıt ol çıkış
	r.GET("/login", controller.Login)
	r.POST("/login", controller.Login)
	r.POST("/signup", controller.SignUp)
	//	r.POST("/exit", controller.Exit)

	//kullanıcı eğitim eklem ve kendi satın aldığı eğitimlere görme
	r.GET("/mycourse", controller.MyCourse)
	r.GET("/addcourse", controller.AddEducation)
	r.POST("/addcourse", controller.AddEducation)
	r.POST("/course", controller.Course)

	//sepet
	r.POST("/order", controller.Addsepet)
	r.GET("/order", controller.Sepet)
	//kullanıcı video işlemleri
	r.POST("/upload", controller.UploadVideo)

	r.GET("/exit", controller.Exit)

	//kullanıcı yorum işlemleri
	r.POST("/addcomment", controller.Addcomment)
	r.ServeFiles("/user/asset/*filepath", http.Dir("user/asset"))
	r.POST("/deletecourse", controller.EduSil)
	r.GET("/edu/eduvideo", controller.Eduvideo)

	//kart işlemleri
	r.GET("/mycard", controller.Mycard)
	r.POST("/addcard", controller.AddCard)
	r.POST("/removeitem", controller.DeleteCard)
	r.POST("/buy", controller.Buy)

	r.GET("/update", controller.EduGuncelle)
	r.POST("/update", controller.EduGuncelle)
	return r
}
