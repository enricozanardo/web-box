package main

import (
	"log"
	"net/http"
	"github.com/onezerobinary/web-box/handler"
	"os"
)

type msg struct {
	Num int
}


const (
	DEFAULT_PORT = "8800"
)


func main() {

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	// Think about that declaration
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler.SignInHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/policy", handler.PolicyHandler)
	http.HandleFunc("/signin", handler.SignInHandler)
	http.HandleFunc("/signup", handler.SignUpHandler)

	//Ajax controller
	http.HandleFunc("/checksignup", handler.CheckSignup)

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}


func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/images/favicon.ico")
}