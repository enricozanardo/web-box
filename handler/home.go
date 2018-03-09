package handler

import (
	"net/http"
	"html/template"
)

var home = template.Must(template.ParseFiles(
	 "templates/_base.html",
	"templates/index.html",
))

func HomeHandler(w http.ResponseWriter, req *http.Request) {

	// get cookie



	home.Execute(w, nil)
}