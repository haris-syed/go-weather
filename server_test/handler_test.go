package server_test

import (
	"context"
	"go-weather/server"
	"testing"

	"google.golang.org/grpc"
)

func flt(f float32) *float32 {
	return &f
}
func TestGRPCHandler(t *testing.T) {
	go server.StartGRPCServer("9000")
	serverAddress := "localhost:9000"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := server.NewWeatherSearchClient(conn)
	var search server.SearchRequest = server.SearchRequest{LocationLongitude: flt(73.1), LocationLatitude: flt(33.6), Temperature: flt(2.2), Humidity: flt(50)}

	_, err = client.Search(context.Background(), &search)
	if err != nil {
		t.Errorf("GRPC request failed")
	}

	//TODO: do some comparison of response and epected data
}
