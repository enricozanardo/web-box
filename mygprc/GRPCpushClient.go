package mygprc

import (
	"golang.org/x/net/context"
	"github.com/goinggo/tracelog"
	"os"
	pb_push "github.com/onezerobinary/push-box/proto"
	"google.golang.org/grpc"
)

const (
	PUSHSERVICEADDRESS = "localhost:1972"    // Development
	//PUSHSERVICEADDRESS = "172.104.230.81:1972" // Staging environment
)

func StartPushGRPCConnection() (connection *grpc.ClientConn){
	// set up connection to the gRPC server
	conn, err := grpc.Dial(PUSHSERVICEADDRESS, grpc.WithInsecure())
	if err != nil {
		tracelog.Errorf(err, "GRPCpushClient", "StartGRPCConnection", "Did not open the connection")
		os.Exit(1)
	}

	return conn
}

func StopPushGRPCConnection(connection *grpc.ClientConn){
	// set up connection to the gRPC server
	err := connection.Close()

	if err != nil {
		tracelog.Errorf(err, "GRPCpushClient", "StopGRPCConnection", "Did not close the connection")
		os.Exit(1)
	}
}

func SendNotifications(info *pb_push.Info) (response *pb_push.PushResponse) {

	conn := StartPushGRPCConnection()
	defer StopPushGRPCConnection(conn)
	// Search into the DB the user
	client := pb_push.NewPushServiceClient(conn)

	resp, err := client.SendNotifications(context.Background(), info)

	if err != nil {
		tracelog.Errorf(err, "GRPCpushClient", "SendNotifications", "Error: Notifications not sent")
		os.Exit(1)
	}

	if resp.Code != 200 || resp.Code == 0 {
		tracelog.Trace("GRPCpushClient", "SendNotifications", "It was not possible to send the Notification")
	}

	return resp
}