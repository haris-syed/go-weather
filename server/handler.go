package server

import (
	"context"
	"fmt"
	"go-weather/search"
	"log"
	"net"
	"os"

	"github.com/affanshahid/configo"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	UnimplementedWeatherSearchServer
}

func StartGRPCServer(port string) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Failed to listen")
	}
	s := Server{}
	grpcServer := grpc.NewServer()
	RegisterWeatherSearchServer(grpcServer, &s)
	reflection.Register(grpcServer)
	fmt.Println("Server started")

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func (s Server) Search(ctx context.Context, input *SearchRequest) (*SearchResponse, error) {
	err := configo.Initialize(
		os.DirFS("../config"),
		configo.WithDeploymentFromEnv("APP_ENV"),
	)
	if err != nil {
		panic(err)
	}

	location := search.Coordinate{Longitude: *input.LocationLongitude, Latitude: *input.LocationLatitude}
	conditions := search.Conditions{Temp: *input.Temperature, Humidity: int(*input.Humidity)}
	data, err := search.FastFetchData(location,
		fmt.Sprintf(
			"https://api.openweathermap.org/data/2.5/find?lat=%f&lon=%f&cnt=50&units=metric&appid=%s",
			location.Latitude,
			location.Longitude,
			configo.MustGetString("apikey")))
	if err != nil {
		panic(err)
	}
	result := search.QuickSearch(conditions, data)
	var response SearchResponse = SearchResponse{}
	for _, val := range result {
		response.Data = append(response.Data, &WeatherData{Location: &WeatherData_Coordinate{Longitude: val.Coord.Longitude, Latitude: val.Coord.Latitude}, Main: &WeatherData_Maininfo{Temp: val.Main.Temp, Tempmin: val.Main.TempMin, Tempmax: val.Main.TempMax, Feelslike: val.Main.FeelsLike, Pressure: int32(val.Main.Pressure), Humidity: int32(val.Main.Humidity), Sealevel: int32(val.Main.SeaLevel), Grndlevel: int32(val.Main.GrndLevel)}})
	}
	return &response, nil

}
