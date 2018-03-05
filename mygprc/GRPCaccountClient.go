package mygprc


import (
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/goinggo/tracelog"
	"os"
)

const (
	ADDRESS = "localhost:1982"
)

func StartGRPCConnection() (connection *grpc.ClientConn){
	// set up connection to the gRPC server
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "StartGRPCConnection", "Did not open the connection")
		os.Exit(1)
	}

	return conn
}

func StopGRPCConnection(connection *grpc.ClientConn){
	// set up connection to the gRPC server
	err := connection.Close()

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "StopGRPCConnection", "Did not close the connection")
		os.Exit(1)
	}
}


func CreateAccount(account *pb_account.Account) (accountResponse pb_account.Response)  {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)
	// Search into the DB the user
	client := pb_account.NewAccountServiceClient(conn)

	resp, err := client.CreateAccount(context.Background(), account)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "CreateAccount", "Error: Account not created")
		os.Exit(1)
	}

	if resp.Code == 200 {
		tracelog.Trace("GRPCaccountClient", "CreateAccount", "A new Account is added to DB")
	} else {
		tracelog.Trace("GRPCaccountClient", "CreateAccount", "It was not possible to add a new Account into DB")
	}

	return *resp
}
