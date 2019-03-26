package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"awesomeProject/car"
)

const port = ":51001"

type server struct{}

func (s server) Build(ctx context.Context, req *car.CreateCarRequest) (*car.CreateCarResponse, error) {
	log.Infof("Building car... Request {%+v} \n", req)

	st := status.New(codes.NotFound, "test")
	st.WithDetails()
	st.Err()
	return &car.CreateCarResponse{
		Identifier: uuid.NewV1().String(),
	}, nil
}

func (s server) List(ctx context.Context, req *car.ListCarRequest) (*car.ListCarResponse, error) {
	return &car.ListCarResponse{
		Cars: []*car.Car{
			{Identifier: uuid.NewV1().String(), Name: "Batmobile"},
			{Identifier: uuid.NewV1().String(), Name: "K2000"},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	car.RegisterFactoryServer(s, &server{})

	log.Infof("gRPC car factory server starting on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
