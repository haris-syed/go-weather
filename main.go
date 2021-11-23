package main

import (
	"fmt"
	"go-weather/server"
	"log"
	"net"
	"os"

	"github.com/affanshahid/configo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// fmt.Println("Hello weather")
	// var london search.Coordinate = search.Coordinate{Latitude: 51.5072, Longitude: -0.1276}
	// londonrain := search.QuickSearch(london, search.Conditions{Temp: 8, Humidity: 60})
	// for _, val := range londonrain {
	// 	fmt.Printf("%+v\n", val.Main.Temp)
	// }
	// fmt.Printf("%+v\n", londonrain)
	// fmt.Println("No. of Results:", len(londonrain))
	err := configo.Initialize(
		os.DirFS("./config"),
		configo.WithDeploymentFromEnv("APP_ENV"),
	)
	if err != nil {
		panic(err)
	}
	port := configo.MustGetString("grpcport")

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Failed to listen")
	}
	s := server.Server{}
	grpcServer := grpc.NewServer()
	server.RegisterWeatherSearchServer(grpcServer, &s)
	reflection.Register(grpcServer)
	fmt.Println("Server started")

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}

}
