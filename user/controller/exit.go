package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Exit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	cookie := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
