package handler

import (
	"net/http"
	"github.com/goinggo/tracelog"
)

func SignOutHandler(w http.ResponseWriter, req *http.Request) {

	// Get the cookie
	sessionCookie, err := req.Cookie("session")

	if err != nil {
		tracelog.Errorf(err, "signout", "SignOutHandler", "SessionCookie not removed.")
	}

	//Remove from the session manager the current sessionCookie
	delete(dbSession, sessionCookie.Value)

	// remove the cookie
	sessionCookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, sessionCookie)

	http.Redirect(w, req, "/", http.StatusSeeOther)

	return
}