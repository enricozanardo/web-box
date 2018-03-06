package mygprc

import (
	"testing"
	"github.com/goinggo/tracelog"
	pb_account "github.com/onezerobinary/db-box/proto/account"
)

func TestCreateAccount(t *testing.T){

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//fakeStatus := pb_account.Status{pb_account.Status_NOTSET}

	username := "Zorro"
	password := "Zirro"

	faketoken := GenerateToken(username, password)

	fakeAccount := pb_account.Account{
		faketoken,
		username,
		password,
		nil,
		nil,
		"Account",
		"2018-01-11",
		"2028-01-10",
	}

	resp:= CreateAccount(&fakeAccount)

	if resp.Code != 200 {
		t.Error("Not possible to create an Account")
	}
}


func TestGetAccountByToken(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//username := "Zorro"
	//password := "Zirro"

	//faketoken := GenerateToken(username, password)

	token := pb_account.Token{"ae30c48577a9e11421ed5434259c3fd88a6c5587"}

	account := GetAccountByToken(&token)

	if account.Token.Token != token.Token {
		t.Errorf("Error in retrieving the account")
	}
}


func TestGetAccountByCredentials(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	username := "Tino"
	password := "Zuccon"

	faketoken := GenerateToken(username, password)

	token := pb_account.Token{faketoken}

	credentials := pb_account.Credentials{username, password, &token}

	account := GetAccountByCredentials(&credentials)

	if account.Token.Token != token.Token {
		t.Errorf("Error in retrieving the account")
	}
}