package mygprc

import (
	"os"
	"github.com/spf13/viper"
	"github.com/goinggo/tracelog"
	"google.golang.org/grpc"
	pb_geo "github.com/onezerobinary/geo-box/proto"
	"context"
)

func StartGeoGRPCConnection() (connection *grpc.ClientConn){

	// Get info from production or local
	EmailServiceAddress := os.Getenv("GEO_SERVICE_ADDRESS")

	if len(EmailServiceAddress) == 0 {
		EmailServiceAddress = viper.GetString("service.geo-box")
		tracelog.Warning("GRPCgeoClient", "StartGRPCConnection", "####### Development #########")
	}

	// set up connection to the gRPC server
	conn, err := grpc.Dial(EmailServiceAddress, grpc.WithInsecure())
	if err != nil {
		tracelog.Errorf(err, "GRPCgeoClient", "StartGRPCConnection", "Did not open the connection")
		os.Exit(1)
	}

	return conn
}

func StopGeoGRPCConnection(connection *grpc.ClientConn){
	// set up connection to the gRPC server
	err := connection.Close()

	if err != nil {
		tracelog.Errorf(err, "GRPCgeoClient", "StopGRPCConnection", "Did not close the connection")
		os.Exit(1)
	}
}

func CalculatePoint(address pb_geo.Address) (point *pb_geo.Point){

	conn := StartPushGRPCConnection()
	defer StopPushGRPCConnection(conn)
	// Get the point providing the address
	client := pb_geo.NewGeoServiceClient(conn)

	point, err := client.GetPoint(context.Background(), &address)

	if err != nil {
		tracelog.Error(err, "CalculatePoint", "It was not possible to get the point")
		return &pb_geo.Point{}
	}

	return point
}


func GetDevices(researchArea pb_geo.ResearchArea) (devices *pb_geo.Devices) {

	conn := StartPushGRPCConnection()
	defer StopPushGRPCConnection(conn)
	// Get the point providing the address
	client := pb_geo.NewGeoServiceClient(conn)

	devices, err :=client.GetDeviceList(context.Background(), &researchArea)

	if err != nil {
		tracelog.Error(err, "GetDevices", "It was not possible to get the devices")
	}

	return devices
}


