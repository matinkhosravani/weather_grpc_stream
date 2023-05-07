package main

import (
	"flag"
	"fmt"
	pb "github.com/matinkhosravani/weather_stream_grpc/proto"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

type weatherServer struct {
	pb.UnimplementedWeatherServiceServer
}

var (
	port = flag.String("port", "50051", "port to listen")
)

func (s *weatherServer) GetUpdates(loc *pb.Location, stream pb.WeatherService_GetUpdatesServer) error {
	for {
		// generate random weather data
		temp := rand.Float64()*10 + 20 // temperature between 20 and 30 degrees Celsius
		humidity := rand.Float64() * 0.2
		windSpeed := rand.Float64()*10 + 5 // wind speed between 5 and 15 km/h

		// create a weather update object
		update := &pb.WeatherUpdate{
			LocationName: loc.Name,
			Temperature:  temp,
			Humidity:     humidity,
			WindSpeed:    windSpeed,
		}

		// send the weather update to the client
		if err := stream.Send(update); err != nil {
			log.Printf("error sending update: %v", err)
			return err
		}

		time.Sleep(2 * time.Second)
	}
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &weatherServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
