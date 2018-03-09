package handler

import (
	"net/http"
	"html/template"
	"github.com/onezerobinary/web-box/model"
)

var policy = template.Must(template.ParseFiles(
	"templates/_base.html",
	 "templates/policy.html",
))

func PolicyHandler(w http.ResponseWriter, req *http.Request) {

	loggedIn := AlreadyLoggedIn(req)

	message := model.MessageLoggedIn{}

	if loggedIn {
		//Redirect to home
		http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
	}

	message.AlreadyLoggedIn = false

	policy.Execute(w, message)

}
