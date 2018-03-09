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
	loggedIn := AlreadyLoggedIn(req)

	message := model.MessageLoggedIn{}

	if !loggedIn {
		//Redirect to home
		http.Redirect(w, req, "/", http.StatusSeeOther)
	} else {
		// is logged in!
		message.AlreadyLoggedIn = true
	}

	dashboard.Execute(w, message)
}


func AlreadyLoggedIn(req *http.Request) bool {

	cookie, err := req.Cookie("session")

	if err != nil {
		return false
	}

	token := dbSession[cookie.Value]

	if token.Token != "" {
		return true
	}

	//TODO: check that the token is good!

	//// Connect to the service
	//conn := api.StartGRPCConnection()
	//defer api.StopGRPCConnection(conn)
	//client := pb.NewAccountServiceClient(conn)
	//// Search into the DB the user
	//account := api.GetAccount(client, &token)
	//
	//// If true already logged in
	//if account.Username != ""  {
	//	//fmt.Println("Already logged in! ", account.Username)
	//	return true
	//}

	return false
}


