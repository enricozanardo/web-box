package handler

import (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
)

var signin = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/signin.html",
))


func SignInHandler(w http.ResponseWriter, req *http.Request) {

	// Get the cookie
	sessionCookie, err := req.Cookie("session")

	if err != nil {
		// Create a new cookie
		sessionID := uuid.NewV4()
		sessionCookie = &http.Cookie{
			Name: "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, sessionCookie)
	}

	// render the page
	signin.Execute(w, nil)
}