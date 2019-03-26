package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"awesomeProject/car"
)

const address = ":51001"

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := car.NewFactoryClient(conn)

	// timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Build(ctx, &car.CreateCarRequest{Name: "Audi A3", Color: car.CreateCarRequest_BLACK, Cooler: true})
	if err != nil {
		log.Fatalf("Building car err: %+v", err)
	}
	log.Printf("Car: %v", resp)
}
