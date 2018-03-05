package mygprc

import (
	"testing"
	"github.com/goinggo/tracelog"
	pb_account "github.com/onezerobinary/db-box/proto/account"
)

func TestCreateAccount(t *testing.T){

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//fakeToken := pb_account.Token{"fff"}
	fakeStatus := pb_account.Status{pb_account.Status_NOTSET}

	fakeAccount := pb_account.Account{
		"168-134-124-1",
		"gino",
		"bari",
		nil,
		&fakeStatus,
		"Account",
		"2018-01-11",
		"2028-01-10",
	}

	resp:= CreateAccount(&fakeAccount)

	if resp.Code != 200 {
		t.Error("Not possible to create an Account")
	}
}
