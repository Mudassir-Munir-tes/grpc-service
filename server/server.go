package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Mudassir-Munir-tes/grpc-service/companypb"
	"github.com/Mudassir-Munir-tes/grpc-service/database"
	"github.com/Mudassir-Munir-tes/grpc-service/models"
	"google.golang.org/grpc"
)

type Server struct {
	companypb.UnimplementedDriverServiceServer
}

func init() {
	database.Connect()
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	srv := grpc.NewServer()

	companypb.RegisterDriverServiceServer(srv, &Server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}

func (srv *Server) InsertDriver(ctx context.Context, request *companypb.DriverRequest) (*companypb.DriverResponse, error) {

	name := request.GetRequest().GetName()
	driver := &models.Driver{Name: name}

	database.Db.Save(driver)

	response := &companypb.DriverResponse{Id: int32(driver.ID), Response: &companypb.Driver{Name: name}}

	return response, nil
}

func (srv *Server) InsertTruck(ctx context.Context, request *companypb.TruckRequest) (*companypb.TruckResponse, error) {

	modelNo, power := request.GetRequest().GetModelNo(), request.GetRequest().GetPower()

	truck := &models.Truck{ModelNo: modelNo, Power: power}

	database.Db.Save(truck)

	response := &companypb.TruckResponse{Id: int32(truck.ID), Response: &companypb.Truck{ModelNo: modelNo, Power: power}}
	return response, nil
}

func (srv *Server) InsertUser(ctx context.Context, request *companypb.UserRequest) (*companypb.UserResponse, error) {
	name := request.GetRequest().GetName()
	fmt.Println(name)

	user := &models.User{Name: name}

	fmt.Println(user)

	database.Db.Save(user)

	fmt.Println(user)

	response := &companypb.UserResponse{Id: int32(user.ID), Response: &companypb.User{Name: name}}

	return response, nil

}
