package handler

import (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
	"fmt"
	"encoding/json"
)

var signup = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/signup.html",
))

func SignUpHandler(w http.ResponseWriter, req *http.Request) {

	// Get the cookie
	sessionCookie, err := req.Cookie("session")

	if err != nil {
		// Create a new cookie
		sessionID := uuid.NewV4()
		sessionCookie = &http.Cookie{
			Name: "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, sessionCookie)
	}

	// render the page
	signup.Execute(w, nil)
}


type Response struct {
	Email string
	Password string
	Policy string
}

func CheckSignup(w http.ResponseWriter, req *http.Request) {

	response := Response{}

	// process form submission
	if req.Method == http.MethodPost {

		usr := req.FormValue("inputEmail")
		psw1 := req.FormValue("inputPassword1")
		psw2 := req.FormValue("inputPassword2")
		check := req.FormValue("check")

		//TODO: Check if the email is already taken
		remove := "aaa@aaa.com"

		if usr == remove {
			// Already taken
			response.Email = usr
		}

		//Check that the password is not empty
		if len(psw1) | len(psw2) == 0 {
			response.Password = "Enter a password."
		}

		//Check that the 2 password match
		if psw1 != psw2 {
			response.Password = "Your passwords did not match. Please re-enter your passwords."
		}

		if len(check) == 0 {
			response.Policy = "Policy must accepted in order to crate an account"
		}

		byteSlice, _ := json.Marshal(response)

		fmt.Fprint(w, string(byteSlice))
	}
}



