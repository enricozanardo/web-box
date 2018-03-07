package handler

import (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
	"github.com/goinggo/tracelog"
)

var home = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))


func HomeHandler(w http.ResponseWriter, req *http.Request) {
	// get cookie
	sessionCookie, err := req.Cookie("session")
	if err != nil {
		sID, err := uuid.NewV4()

		if err != nil {
			tracelog.Errorf(err, "home", "HomeHandler", "Not able to generate the uuid")
		}

		sessionCookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, sessionCookie)
	}

	home.Execute(w, nil)
}