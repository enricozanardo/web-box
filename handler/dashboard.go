package handler

import (
	"net/http"
	"html/template"
	"github.com/onezerobinary/web-box/model"
	"encoding/json"
	"fmt"
	"github.com/onezerobinary/web-box/mygprc"
	pb_push "github.com/onezerobinary/push-box/proto"
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"github.com/goinggo/tracelog"
	"time"
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

	//TODO: Update the map lat and ltg based on the user login

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


func PushHandler(w http.ResponseWriter, req *http.Request){

	message := model.EmergencyMessage{}

	if req.Method == http.MethodPost {

		notification := pb_push.Info{}

		//Retrieve all the data from the form
		emergencyAddress := req.FormValue("inputAddress")
		emergencyNumber := req.FormValue("inputNumber")
		emergencyPostalCode := req.FormValue("inputPostalCode")
		emergencyPlace := req.FormValue("inputPlace")

		//Check the data
		if len(emergencyAddress) == 0 {
			message.EmergencyAddressMessage  = "Enter an address."
		}

		if len(emergencyNumber) == 0 {
			message.EmergencyAddressMessage  = "Enter an address number."
		}

		if len(emergencyPostalCode) == 0 {
			message.EmergencyAddressMessage  = "Enter a postal code."
		}

		if len(emergencyPlace) == 0 {
			message.EmergencyAddressMessage  = "Enter a place."
		}

		//use geo-box
		address := pb_geo.Address{}
		address.Address = emergencyAddress
		address.AddressNumber = emergencyNumber
		address.PostalCode = emergencyPostalCode
		address.Place = emergencyPlace

		point := mygprc.CalculatePoint(address)

		//From float32 to string
		emergencyLat := fmt.Sprintf("%f", point.Latitude)
		emergencyLng := fmt.Sprintf("%f", point.Longitude)

		//use geo-box
		//Collect the nearest devices
		researchArea := pb_geo.ResearchArea{}
		researchArea.Point = point
		researchArea.Precision = 5

		devices := mygprc.GetDevices(researchArea)

		//token := pb_account.Token{"d0a1a743194ff28f049f47b9b69c51563c2cfadf"}
		//
		//fakeAccount := mygprc.GetAccountByToken(&token)

		for _, device := range devices.Expopushtoken {
			notification.DeviceTokens = append(notification.DeviceTokens, device)
		}

		//trasform the time in the right format
		emergencyTime := time.Now()
		// Set the layout that are needed into the DB
		layout := "2006-01-02T15:04:05.000Z"
		etString := string(emergencyTime.Format(layout))

		//Generate the emergency
		emegency := pb_push.Emergency{}
		emegency.Address = emergencyAddress
		emegency.AddressNumber = emergencyNumber
		emegency.PostalCode = emergencyPostalCode
		emegency.Place = emergencyPlace
		emegency.Lat = emergencyLat
		emegency.Lng = emergencyLng
		emegency.Time = etString
		emegency.IsActive = true

		//Send the Notification
		notification.Emergency = &emegency

		resp := mygprc.SendNotifications(&notification)

		if resp.Code != 200 || resp.Code != 0 {
			warn := "Possible error in sending notifications" + string(resp.Code)
			tracelog.Warning("dashboard", "PushHandler", warn)
			//Notify the user
			message.EmergencySuccessMessage = "Possible error in sending notifications"
		} else {
			message.EmergencySuccessMessage = "Notifications successfully sent"
		}

		//TODO: Store the Emergency into data into the DB

	}

	// send back the errors!
	byteSlice, _ := json.Marshal(message)

	// clean the message
	message = model.EmergencyMessage{}

	fmt.Fprint(w, string(byteSlice))
}
