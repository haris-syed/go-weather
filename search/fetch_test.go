package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFetchData(t *testing.T) {

	testData, err := ioutil.ReadFile("../fetchtestdata.json")
	if err != nil {
		panic(err)
	}
	var expected apiresponse
	if err := json.Unmarshal(testData, &expected); err != nil {
		panic(err)
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(testData))
	}))
	defer svr.Close()

	api := OpenWatherApi{Url: svr.URL}
	response := api.fetchData(Coordinate{0, 0})

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("Response not matched")
	}
}
