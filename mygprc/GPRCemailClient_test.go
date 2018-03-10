package mygprc

import (
	"testing"
	pb_email "github.com/onezerobinary/email-box/proto"
)

func TestSendEmail(t *testing.T) {

	fakeRecipient := pb_email.Recipient{"enrico.zanardo101@gmail.com", "ABC123456", 0}

	response := SendEmail(&fakeRecipient)

	if response.Code != 200 {
		t.Errorf("Error in sending the email")
	}
}
