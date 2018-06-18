package main

import (
	"log"
	"net/http"
	"github.com/onezerobinary/web-box/handler"
	"os"
	"github.com/goinggo/tracelog"
	"github.com/spf13/viper"
)

type msg struct {
	Num int
}

const (
	DEFAULT_PORT = "8800"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	//development environment
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "main", "main", "Error reading config file")
	}

	tracelog.Warning("main", "main", "Using config file")

	// Think about that declaration
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler.SignInHandler)
	http.HandleFunc("/confirm", handler.ConfirmHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/dashboard", handler.DashboardHandler)
	http.HandleFunc("/policy", handler.PolicyHandler)
	http.HandleFunc("/signin", handler.SignInHandler)
	http.HandleFunc("/signup", handler.SignUpHandler)
	http.HandleFunc("/signout", handler.SignOutHandler)
	http.HandleFunc("/disabled", handler.DisabledHandler)
	http.HandleFunc("/suspended", handler.DisabledHandler)
	http.HandleFunc("/revoked", handler.DisabledHandler)

	//Ajax controller
	http.HandleFunc("/checksignup", handler.CheckSignup)
	http.HandleFunc("/checksignin", handler.CheckSignin)
	http.HandleFunc("/push", handler.PushHandler)

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}


func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/images/favicon.ico")
}