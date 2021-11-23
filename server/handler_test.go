package server

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
)

func flt(f float32) *float32 {
	return &f
}
func TestGRPCHandler(t *testing.T) {
	go StartGRPCServer("9000")
	serverAddress := "localhost:9000"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := NewWeatherSearchClient(conn)
	var search SearchRequest = SearchRequest{LocationLongitude: flt(73.1), LocationLatitude: flt(33.6), Temperature: flt(2.2), Humidity: flt(50)}
	response, _ := client.Search(context.Background(), &search)
	//TODO: do some comparsion with an expected value
	fmt.Println(response.Data)
}
