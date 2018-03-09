package handler

import (
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/goinggo/tracelog"
	"github.com/onezerobinary/web-box/model"
	"github.com/onezerobinary/web-box/mygprc"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"fmt"
	"encoding/json"
	"html/template"
)

var signin = template.Must(template.ParseFiles(
	 "templates/_base.html", "templates/signin.html",
))

var dbSession = map[string]pb_account.Token{} // sessionID, userID


func SignInHandler(w http.ResponseWriter, req *http.Request) {

	// render the page

	signin.Execute(w, nil)
}


func GetSessionCookie(w http.ResponseWriter, req *http.Request)(sessionCookie *http.Cookie) {
	// Get the cookie
	sessionCookie, err := req.Cookie("session")

	if err != nil {
		// Create a new cookie
		sessionID, err := uuid.NewV4()

		if err != nil {
			tracelog.Errorf(err, "signin", "GetSessionCookie", "Not able to generate the uuid")
		}

		sessionCookie = &http.Cookie{
			Name: "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, sessionCookie)
	}

	return sessionCookie
}


func CheckSignin(w http.ResponseWriter, req *http.Request) {

	message := model.Message{}

	// process form submission
	if req.Method == http.MethodPost {

		email := req.FormValue("inputEmail")
		password := req.FormValue("inputPassword")

		userToken := mygprc.GenerateToken(email, password)

		token := pb_account.Token{userToken}

		credentials := pb_account.Credentials{email, password, &token}

		account := mygprc.GetAccountByCredentials(&credentials)


		if account.Token.Token != token.Token {
			//User not present into the system -> Signup ?
			message.LoginMessage = "User not present into the system, please - <a href=\"/signup\">singup</a>"
		}

		if account.Token.Token == token.Token {
			//Only ENABLED account are allowed to access to the system
			if account.Status.Status != pb_account.Status_ENABLED {
				message.LoginMessage = "Account not allowed to access: " + account.Status.Status.String()
			}
		}

		//Check that the email is not empty
		if len(email) == 0 {
			message.EmailMessage = "Email not provided"
		}

		//Check that the password is not empty
		if len(password) == 0 {
			message.PasswordMessage = "Password not provided"
		}

		if (model.Message{}) == message  {
			//  Associate the user token to the session
			cookie := GetSessionCookie(w, req)
			dbSession[cookie.Value] = token

			tracelog.Trace("signin", "CheckSignin", "Token added to Session")
			//TODO: perform the login!


		}

		// send back the errors!
		byteSlice, _ := json.Marshal(message)

		// clean the message
		message = model.Message{}

		fmt.Fprint(w, string(byteSlice))
	}
}