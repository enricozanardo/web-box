package handler

import (
	"net/http"
	"html/template"
)

var policy = template.Must(template.ParseFiles(
	"templates/_base.html",
	 "templates/policy.html",
))

func PolicyHandler(w http.ResponseWriter, req *http.Request) {

	policy.Execute(w, nil)
}
