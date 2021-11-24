package main

import (
	"fmt"
	"go-weather/server"
	"os"

	"github.com/affanshahid/configo"
)

func main() {
	err := configo.Initialize(
		os.DirFS("./config"),
		configo.WithDeploymentFromEnv("APP_ENV"),
	)
	if err != nil {
		panic(err)
	}
	port := configo.MustGetString("grpcport")
	server.StartGRPCServer(port)
	fmt.Println("Hello after server")

}
