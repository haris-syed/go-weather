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

func TestSearch(t *testing.T) {
	testData, err := ioutil.ReadFile("../fetchtestdata.json")
	if err != nil {
		panic(err)
	}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(testData))
	}))
	defer svr.Close()

	api := OpenWatherApi{Url: svr.URL}
	result := api.QuickSearch(Coordinate{0, 0}, Conditions{Temp: 2.20, Humidity: 50})
	expectedfile, err := ioutil.ReadFile("../searchtestdata.json")
	if err != nil {
		panic(err)
	}
	var expected []Weatherdata
	if err := json.Unmarshal(expectedfile, &expected); err != nil {
		panic(err)
	}

	if reflect.DeepEqual(expected, result) {
		t.Error("Serach result not matched")
	}

}
