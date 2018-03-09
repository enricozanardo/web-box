package handler

import (
	"net/http"
	"html/template"
	"github.com/onezerobinary/web-box/model"
)

var dashboard = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/dashboard.html",
))


func DashboardHandler(w http.ResponseWriter, req *http.Request) {

	//TODO: check if authenticated

	message := model.MessageLoggedIn{}

	// is logged in!
	message.AlreadyLoggedIn = true


	dashboard.Execute(w, message)
}


