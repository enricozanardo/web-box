package handler

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/onezerobinary/web-box/model"
	"github.com/goinggo/tracelog"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	pb_email "github.com/onezerobinary/email-box/proto"
	"github.com/onezerobinary/web-box/mygprc"
	"time"
	"github.com/pkg/errors"
	"html/template"
)

var signup = template.Must(template.ParseFiles(
	"templates/_base.html",
	 "templates/signup.html",
))

func SignUpHandler(w http.ResponseWriter, req *http.Request) {

	loggedIn := AlreadyLoggedIn(req)

	message := model.MessageLoggedIn{}

	if loggedIn {
		//Redirect to home
		http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
	}

	message.AlreadyLoggedIn = false

	signup.Execute(w, message)
}


func CheckSignup(w http.ResponseWriter, req *http.Request) {

	message := model.Message{}

	// process form submission
	if req.Method == http.MethodPost {

		usr := req.FormValue("inputEmail")
		psw1 := req.FormValue("inputPassword1")
		psw2 := req.FormValue("inputPassword2")
		check := req.FormValue("check")

		//Check if the email is already taken "aaa@bbb.ccc"
		dbEmail := pb_account.Email{usr}

		response := mygprc.CheckEmail(&dbEmail)

		if response.Code == 200 && response.Token != nil {
			tracelog.Trace("signup", "CheckSignup", "Email already taken")
			message.EmailMessage = "Username already taken: " + usr + " - <a href=\"/signin\">singin</a>?"
		}

		//Check that the password is not empty
		if len(psw1) | len(psw2) == 0 {
			message.PasswordMessage = "Enter a password."
		}

		//Check that the 2 password match
		if psw1 != psw2 {
			message.PasswordMessage = "Your passwords did not match. Please re-enter your passwords."
		}

		if len(check) == 0 {
			message.PolicyMessage = "Policy must accepted in order to be able to create a new account"
		}

		if (model.Message{}) == message  {
			//create the new user!
			err := createNewAccount(usr, psw1)

			if err != nil {
				message.LoginMessage = err.Error()
			} else {
				//if ok then send the email
				message.Allowed = true
				message.LoginMessage = "Well done! Your account has been made, please verify it by clicking the activation link that has been send to your email"

				userToken := mygprc.GenerateToken(usr, psw1)
				recipient := pb_email.Recipient{usr,userToken, 0}

				response := mygprc.SendEmail(&recipient)

				if response.Code != 200 {
					message.LoginMessage = "It was not possible to send the email"
				} else {
					message.LoginMessage = "Well done! Your account has been made, please verify it by clicking the activation link that has been send to your email"
				}
			}
		}

		// send back the errors!
		byteSlice, _ := json.Marshal(message)

		// clean the message
		message = model.Message{}

		fmt.Fprint(w, string(byteSlice))
	}
}


func createNewAccount(username, password string) (err error) {
	fmt.Printf("OK! %v %v", username, password)

	userToken := mygprc.GenerateToken(username, password)

	token := pb_account.Token{userToken}
	status := pb_account.Status{pb_account.Status_NOTSET}

	created := time.Now()
	expiration := created.Add(time.Duration(24*time.Hour))

	// Set the layout that are needed into the DB
	layout := "2006-01-02T15:04:05.000Z"
	c := string(created.Format(layout))
	e := string(expiration.Format(layout))

	account := pb_account.Account{
		userToken,
		username,
		password,
		&token,
		&status,
		"Account",
		c,
		e,
	}

	resp := mygprc.CreateAccount(&account)

	if resp.Code != 200 {
		return errors.New("An error is occurred during the creation of the account")
	}

	return nil
}


