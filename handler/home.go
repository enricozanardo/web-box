package handler

import (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/onezerobinary/web-box/model"
)

var home = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))


func HomeHandler(w http.ResponseWriter, req *http.Request) {
	// get cookie
	sessionCookie, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		sessionCookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, sessionCookie)
	}

	warning := "Ciao"
	m := model.Message{Warning: warning}

	fmt.Println("message: ", m)
	//home.Execute(w, m)

	home.Execute(w, nil)
}