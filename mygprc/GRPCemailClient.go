package mygprc

import (
	"golang.org/x/net/context"
	"github.com/goinggo/tracelog"
	"os"
	pb_email "github.com/onezerobinary/email-box/proto"
	"google.golang.org/grpc"
)

const (
	//EMAILSERVICEADDRESS = "localhost:1976"    // Development
	//EMAILSERVICEADDRESS = "172.104.230.81:1976" // Staging environment
	EMAILSERVICEADDRESS = "46.101.149.16:1976" // Staging environment
)

func StartEmailGRPCConnection() (connection *grpc.ClientConn){
	// set up connection to the gRPC server
	conn, err := grpc.Dial(EMAILSERVICEADDRESS, grpc.WithInsecure())
	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "StartGRPCConnection", "Did not open the connection")
		os.Exit(1)
	}

	return conn
}

func StopEmailGRPCConnection(connection *grpc.ClientConn){
	// set up connection to the gRPC server
	err := connection.Close()

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "StopGRPCConnection", "Did not close the connection")
		os.Exit(1)
	}
}

func SendEmail(recipient *pb_email.Recipient) (response *pb_email.EmailResponse) {

	conn := StartEmailGRPCConnection()
	defer StopEmailGRPCConnection(conn)
	// Search into the DB the user
	client := pb_email.NewEmailServiceClient(conn)

	resp, err := client.SendEmail(context.Background(), recipient)

	if err != nil {
		tracelog.Errorf(err, "GRPCemailClient", "SendEmail", "Error: Email not sent")
		os.Exit(1)
	}

	if resp.Code != 200 {
		tracelog.Trace("GRPCemailClient", "SendEmail", "It was not possible to send the email")
	}

	return resp
}