package mygprc

import (
	"golang.org/x/net/context"
	"github.com/goinggo/tracelog"
	"os"
	"crypto/sha1"
	"encoding/hex"
	"io"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"google.golang.org/grpc"
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

func GenerateToken(username string, password string) (token string){
	// Create the hash sha1 of the username and password
	h1 := sha1.New()
	io.WriteString(h1, username)
	io.WriteString(h1, password)

	token = hex.EncodeToString(h1.Sum(nil))

	tracelog.Completed("database","GenerateToken")
	return token
}

func CreateAccount(account *pb_account.Account) (response *pb_account.Response)  {

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

	return resp
}

func GetAccountByCredentials(credentials *pb_account.Credentials)(account *pb_account.Account) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_account.NewAccountServiceClient(conn)

	account, err := client.GetAccountByCredentials(context.Background(), credentials)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "GetAccountByCredentials", "Error: It was not possible to retrieve the account")
		os.Exit(1)
	}

	tracelog.Trace("GRPCaccountClient", "GetAccountByCredentials", "An account is retrieved")

	return account
}

// Get an Account given the Token
func GetAccountByToken(token *pb_account.Token)(account *pb_account.Account) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_account.NewAccountServiceClient(conn)

	account, err := client.GetAccountByToken(context.Background(), token)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "GetAccountByToken", "Error: It was not possible to retrieve the account")
		os.Exit(1)
	}

	tracelog.Trace("GRPCaccountClient", "GetAccountByToken", "An account is retrieved")

	return account
}

// Update an Account given the updated Account
func UpdateAccount (account *pb_account.Account )(response *pb_account.Response) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)
	// Search into the DB the user
	client := pb_account.NewAccountServiceClient(conn)

	resp, err := client.UpdateAccount(context.Background(), account)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "UpdateAccount", "Error: Account not updated")
		os.Exit(1)
	}

	if resp.Code == 200 {
		tracelog.Trace("GRPCaccountClient", "UpdateAccount", "Account is updated into DB")
	} else {
		tracelog.Trace("GRPCaccountClient", "UpdateAccount", "It was not possible to update the Account into DB")
	}

	return resp
}

// Delete an Account given the Token
func DeleteAccount (token *pb_account.Token)(response *pb_account.Response) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)
	// Search into the DB the user
	client := pb_account.NewAccountServiceClient(conn)

	resp, err := client.DeleteAccount(context.Background(), token)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "DeleteAccount", "Error: It is not possible to delete the account")
		os.Exit(1)
	}

	if resp.Code == 200 {
		tracelog.Trace("GRPCaccountClient", "DeleteAccount", "Account deleted from DB")
	} else {
		tracelog.Trace("GRPCaccountClient", "DeleteAccount", "It was not possible to delete the Account from the DB")
	}

	return resp
}

// Check if an email address is already used
func CheckEmail (email *pb_account.Email)(response *pb_account.Response) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)
	// Search into the DB the user
	client := pb_account.NewAccountServiceClient(conn)

	resp, err := client.CheckEmail(context.Background(), email)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "CheckEmail", "Error: It is not possible to check the email")
		os.Exit(1)
	}

	if resp.Code == 200 {
		tracelog.Trace("GRPCaccountClient", "CheckEmail", "Email successfully checked")
	} else {
		tracelog.Trace("GRPCaccountClient", "CheckEmail", "It was not possible to check the given email into the DB")
	}

	return resp
}

// Get the Status of an account given the Token
func GetAccountStatus (token *pb_account.Token) (status *pb_account.Status) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_account.NewAccountServiceClient(conn)

	status, err := client.GetAccountStatus(context.Background(), token)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "GetAccountStatus", "Error: It was not possible to retrieve the account status")
		os.Exit(1)
	}

	tracelog.Trace("GRPCaccountClient", "GetAccountStatus", "Account status is retrieved")

	return status
}

// Set the Status of an account given the Updated Status
func SetAccountStatus (updateStatus *pb_account.UpdateStatus)(response *pb_account.Response) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)
	// Search into the DB the user
	client := pb_account.NewAccountServiceClient(conn)

	resp, err := client.SetAccountStatus(context.Background(), updateStatus)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "SetAccountStatus", "Error: It is not possible to set the account status")
		os.Exit(1)
	}

	if resp.Code == 200 {
		tracelog.Trace("GRPCaccountClient", "SetAccountStatus", "Account status successfully updated")
	} else {
		tracelog.Trace("GRPCaccountClient", "SetAccountStatus", "It was not possible to update the account status into the DB")
	}

	return resp
}

// Get all the accounts based on a specific Status
func GetAccountsByStatus (status *pb_account.Status)(accounts *pb_account.Accounts ) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_account.NewAccountServiceClient(conn)

	accounts, err := client.GetAccountsByStatus(context.Background(), status)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "GetAccountsByStatus", "Error: It was not possible to retrieve the accounts")
		os.Exit(1)
	}

	tracelog.Trace("GRPCaccountClient", "GetAccountsByStatus", "Accounts are successfully retrieved")

	return accounts
}

//Get the account collection
func GetAccounts (empty *pb_account.Empty)(accounts *pb_account.Accounts) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_account.NewAccountServiceClient(conn)

	accounts, err := client.GetAccounts(context.Background(), empty)

	if err != nil {
		tracelog.Errorf(err, "GRPCaccountClient", "GetAccounts", "Error: It was not possible to retrieve the accounts")
		os.Exit(1)
	}

	tracelog.Trace("GRPCaccountClient", "GetAccounts", "Accounts are successfully retrieved")

	return accounts
}



