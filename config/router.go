package config

import (
	"net/http"
	"new_CodingTime/user/controller"

	"github.com/julienschmidt/httprouter"
)

func Router_Config() *httprouter.Router {
	r := httprouter.New()

	//visiters
	r.GET("/", controller.Index_html)
	r.GET("/login", controller.Login_html)
	r.GET("/blog", controller.Blog_html)

	//LOGİN-SİGNUP
	r.POST("/uslogin", controller.UsLogin)
	r.POST("/ussignup", controller.UsSignUp)

	//education
	//-->/edu/*FILEPATH
	r.GET("/edu/databases", controller.DatabesesIndex)

	//	r.GET("/edu/backend", controller.BackEndIndex)
	r.GET("/edu/etichalhack", controller.EtichalIndex)
	//öğrenim içeriğim
	r.GET("/edu/addcourse", controller.AddCourse)
	//	r.GET("/edu/content", controller.Content)

	//blog
	//add blog
	r.POST("/edu/blog/addblog", controller.Addblog)
	//	r.GET("/deleteblog", controller.DeleteBlog)
	//	r.GET("/updateblog", controller.UpdateBlog)

	//video
	//add video
	r.POST("/edu/eduadd", controller.AddEduCategory)
	//	r.GET("/deletevideo", controller.DeleteVideo)

	//myvideos
	//	r.GET("/myvieos", controller.MyVideos)

	//-->TEACHER
	//-->/edu/admin
	r.GET("/edu/admin", controller.TeacherAdmin)
	//blogAdmin
	//	r.GET("/blog/admin", controller.BlogAdmin)
	//ADMİN DEASHBOARD
	//

	r.POST("/sepeteekle", controller.AddCard)
	r.ServeFiles("/user/asset/*filepath", http.Dir("user/asset"))

	return r
}
