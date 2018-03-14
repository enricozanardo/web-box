package mygprc

import (
	"testing"
	"github.com/goinggo/tracelog"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"fmt"
)

func TestCreateAccount(t *testing.T){

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	username := "enrico@enrico.com"
	password := "enrico"

	faketoken := GenerateToken(username, password)

	fmt.Println(faketoken)

	token := pb_account.Token{faketoken}
	status := pb_account.Status{pb_account.Status_NOTSET}

	fakeAccount := pb_account.Account{
		faketoken,
		username,
		password,
		&token,
		&status,
		"Account",
		"2018-01-11",
		"2028-01-10",
		nil,
	}

	resp:= CreateAccount(&fakeAccount)

	if resp.Code != 200 {
		t.Error("Not possible to create an Account")
	}
}

func TestGetAccountByToken(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	username := "enrico@enrico.com"
	password := "enrico"

	faketoken := GenerateToken(username, password)

	token := pb_account.Token{faketoken}

	account := GetAccountByToken(&token)

	if account.Token.Token != token.Token {
		t.Errorf("Error in retrieving the account")
	}
}

func TestGetAccountByCredentials(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	username := "enrico@enrico.com"
	password := "enrico"

	faketoken := GenerateToken(username, password)

	token := pb_account.Token{faketoken}

	credentials := pb_account.Credentials{username, password, &token}

	account := GetAccountByCredentials(&credentials)

	if account.Token.Token != token.Token {
		t.Errorf("Error in retrieving the account")
	}
}

func TestUpdateAccount(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	username := "enrico@enrico.com"
	password := "enrico"

	faketoken := GenerateToken(username, password)

	token := pb_account.Token{faketoken}
	status := pb_account.Status{pb_account.Status_ENABLED}

	fakeAccount := pb_account.Account{
		faketoken,
		username,
		password,
		&token,
		&status,
		"Account",
		"2018-01-11",
		"2058-01-10",
		nil,
	}

	response := UpdateAccount(&fakeAccount)

	if response.Code != 200 {
		t.Errorf("Error to update the account")
	}
}

func TestCheckEmail(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	fakeEmail := pb_account.Email{"enrico@enrico.com"}

	response := CheckEmail(&fakeEmail)

	if response.Token == nil {
		tracelog.Warning("", "checkemail", "No email are found")
	}

	if response.Code != 200 {
		t.Errorf("Error to check the email")
	}
}
func TestSetAccountStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	faketoken := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}
	newStatus := pb_account.Status{pb_account.Status_ENABLED}

	updateStatus := pb_account.UpdateStatus{&faketoken, &newStatus}

	response := SetAccountStatus(&updateStatus)

	if response.Code != 200 {
		t.Errorf("Error to update the account status")
	}
}

func TestGetAccountStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	faketoken := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}

	status := GetAccountStatus(&faketoken)

	if status.Status == pb_account.Status_NOTSET {
		t.Errorf("Error to retrieve the account status")
	}

}

func TestGetAccountsByStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	fakeStatus := pb_account.Status{pb_account.Status_ENABLED}

	accounts :=GetAccountsByStatus(&fakeStatus)

	if accounts == nil {
		t.Errorf("Empty accounts with the given status are retreived")
	}

	if len(accounts.Accounts) == 0 {
		t.Errorf("Zero accounts with the given status are retreived")
	}
}

func TestGetAccounts(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	empty := pb_account.Empty{}

	accounts := GetAccounts(&empty)

	if accounts == nil {
		t.Errorf("Empty accounts are retreived")
	}

	if len(accounts.Accounts) == 0 {
		t.Errorf("Zero accounts are retreived")
	}
}

func TestDeleteAccount(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	username := "enrico@enrico.com"
	password := "enrico"

	faketoken := GenerateToken(username, password)

	token := pb_account.Token{faketoken}

	response := DeleteAccount(&token)

	if response.Code != 200 {
		t.Errorf("Error to delete the account")
	}
}

