package handler

import (
	"html/template"
	"net/http"
)

var disabled = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/disabled.html",
))


func DisabledHandler(w http.ResponseWriter, req *http.Request) {

	loggedIn := AlreadyLoggedIn(req)

	if loggedIn {
		http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
		return
	}

	disabled.Execute(w, nil)
}