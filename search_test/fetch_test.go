package search

import (
	"encoding/json"
	"errors"
	"go-weather/search"
	"io/ioutil"
	"reflect"
	"testing"
)

type mockWeatherApiClient struct {
}

var mockFetchData func(location search.Coordinate, url string) ([]search.Weatherdata, error)

func (owac mockWeatherApiClient) FetchData(location search.Coordinate, url string) ([]search.Weatherdata, error) {
	return mockFetchData(location, url)
}

func TestFastFetchData(t *testing.T) {

	var mockclient search.WeatherApiInterface = mockWeatherApiClient{}

	testData, err := ioutil.ReadFile("../test_resources/fastfetchtestdata.json")
	if err != nil {
		panic(err)
	}
	var expected []search.Weatherdata
	if err := json.Unmarshal(testData, &expected); err != nil {
		panic(err)
	}

	//test normal case
	mockFetchData = func(location search.Coordinate, url string) ([]search.Weatherdata, error) {
		testData, err := ioutil.ReadFile("../test_resources/fastfetchtestdata.json")
		if err != nil {
			panic(err)
		}
		var res []search.Weatherdata
		if err := json.Unmarshal(testData, &res); err != nil {
			panic(err)
		}
		return res, nil
	}

	data, err := search.FastFetchData(search.Coordinate{0, 0}, "dummy", mockclient)
	if err != nil {
		t.Errorf(err.Error())
	}
	if err != nil || !reflect.DeepEqual(expected, data) {
		t.Errorf("Response not matched")
	}

	//test error case
	mockFetchData = func(location search.Coordinate, url string) ([]search.Weatherdata, error) {
		return nil, errors.New("something wrong")
	}
	_, err = search.FastFetchData(search.Coordinate{0, 0}, "dummy", mockclient)

	if err == nil {
		t.Errorf("Expected error but got result")
	}

}
