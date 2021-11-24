package search_test

import (
	"encoding/json"
	"fmt"
	"go-weather/search"
	"io/ioutil"
	"testing"
)

func compare(a, b []search.Weatherdata) bool {
	for _, val := range a {
		var found bool = false
		for _, val2 := range b {
			if val.ID == val2.ID {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func TestQuickSearch(t *testing.T) {
	expectedfile, err := ioutil.ReadFile("../test_resources/searchtestdata.json")
	if err != nil {
		panic(err)
	}
	var expected []search.Weatherdata
	if err := json.Unmarshal(expectedfile, &expected); err != nil {
		panic(err)
	}

	searchDatafile, err := ioutil.ReadFile("../test_resources/fastfetchtestdata.json")
	if err != nil {
		panic(err)
	}
	var searchData []search.Weatherdata
	if err := json.Unmarshal(searchDatafile, &searchData); err != nil {
		panic(err)
	}

	result := search.QuickSearch(search.Conditions{Temp: 2.20, Humidity: 50}, searchData)

	r, _ := json.Marshal(result)
	fmt.Println(string(r))

	if !compare(result, expected) {
		t.Errorf("The results were not equal")
	}

}
