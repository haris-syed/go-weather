package search

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

type mockWeatherClient struct {
}

var mockFetchData func(location Coordinate, url string) *apiresponse

func (wc mockWeatherClient) fetchData(location Coordinate, url string) *apiresponse {
	return mockFetchData(location, url)
}

func TestFastFetchData(t *testing.T) {
	testData, err := ioutil.ReadFile("../test_resources/fastfetchtestdata.json")
	if err != nil {
		panic(err)
	}
	var expected []Weatherdata
	if err := json.Unmarshal(testData, &expected); err != nil {
		panic(err)
	}

	openWeatherClient = mockWeatherClient{}

	//test normal case
	mockFetchData = func(location Coordinate, url string) *apiresponse {
		testData, err := ioutil.ReadFile("../test_resources/fetchtestdata.json")
		if err != nil {
			panic(err)
		}
		var res *apiresponse
		if err := json.Unmarshal(testData, &res); err != nil {
			panic(err)
		}
		return res
	}

	data, err := FastFetchData(Coordinate{0, 0}, "dummy")

	if err != nil || !reflect.DeepEqual(expected, data) {
		t.Errorf("Response not matched")
	}

	//test error case
	mockFetchData = func(location Coordinate, url string) *apiresponse {
		return nil
	}
	_, err = FastFetchData(Coordinate{0, 0}, "dummy")

	if err == nil {
		t.Errorf("Expected error but got result")
	}

}
