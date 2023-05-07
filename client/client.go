package main

import (
	"context"
	pb "github.com/matinkhosravani/weather_stream_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	cl := pb.NewWeatherServiceClient(conn)

	// create a location object for New York
	loc := &pb.Location{
		Name:      "New York",
		Latitude:  40.7128,
		Longitude: -74.0060,
	}

	// call the GetUpdates method with the location object
	stream, err := cl.GetUpdates(context.Background(), loc)
	if err != nil {
		log.Fatalf("error calling GetUpdates: %v", err)
	}

	// receive and print the weather updates
	for {
		update, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving update: %v", err)
		}
		log.Printf("Received update for %s: temperature=%f, humidity=%f, wind_speed=%f",
			update.LocationName, update.Temperature, update.Humidity, update.WindSpeed)
	}
}
