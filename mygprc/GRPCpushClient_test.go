package mygprc

import (
	"testing"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	pb_push "github.com/onezerobinary/push-box/proto"
	"fmt"
	"github.com/goinggo/tracelog"
	"github.com/onezerobinary/push-box/mygrpc"
)

func TestSendNotifications(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//Fake emergency
	fakeEmegency := pb_push.Emergency{}
	fakeEmegency.Address = "Via Roma"
	fakeEmegency.AddressNumber = "42"
	fakeEmegency.PostalCode = "39100"
	fakeEmegency.Place = "Bolzano"
	fakeEmegency.Lat = "46.4894107"
	fakeEmegency.Lng = "11.3208888"
	fakeEmegency.Time = "2018-03-21T09:47:42.140Z"

	info := pb_push.Info{}
	info.Emergency = &fakeEmegency

	//token := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}
	token := pb_account.Token{"46a249c795cda18c1d8143a781871e1e95d2e011"}

	fakeAccount := mygprc.GetAccountByToken(&token)

	//if err != nil {
	//	tracelog.Error(err, "GRPCpushClient", "TestSendNotification")
	//}

	for _, device := range fakeAccount.Expopushtoken {
		info.DeviceTokens = append(info.DeviceTokens, device)
	}

	status := SendNotifications(&info)

	//if err != nil {
	//	fmt.Println("err: ", err)
	//}

	fmt.Println("code: ", &status)
}
