package mygprc

import (
	"testing"
	pb_email "github.com/onezerobinary/email-box/proto"
	"github.com/goinggo/tracelog"
)


func TestSendEmail(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeRecipient := pb_email.Recipient{"enrico.zanardo101@gmail.com", "ABC123456", 0}

	response := SendEmail(&fakeRecipient)

	if response.Code != 200 {
		t.Errorf("Error in sending the email")
	}
}
