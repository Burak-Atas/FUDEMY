package helpers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("123123"))

func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := store.Get(r, "alert-go")
	if err != nil {
		fmt.Println(err)
		return err
	}
	session.AddFlash(message)

	return session.Save(r, w)
}
