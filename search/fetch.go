package search

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Coordinate struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

type weather struct {
	Id int `json:"id"`
}

type mainInfo struct {
	Temp      float32 `json:"temp"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	FeelsLike float32 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}
type Weatherdata struct {
	ID    int        `json:"id"`
	Name  string     `json:"name"`
	Coord Coordinate `json:"coord"`
	Main  mainInfo   `json:"main,omitempty"`
	Rain  *struct {
		Vol1h float32 `json:"1h"`
		Vol3h float32 `json:"3h"`
	} `json:"rain,omitempty"`
	Snow *struct {
		Vol1h float32 `json:"1h"`
		Vol3h float32 `json:"3h"`
	} `json:"snow,omitempty"`
	Clouds *struct {
		Percentage int `json:"all"`
	} `json:"clouds"`
	Weather []weather `json:"weather"`
}

type apiresponse struct {
	Code  string        `json:"cod"`
	Count int           `json:"cnt"`
	Data  []Weatherdata `json:"list"`
}

type OpenWatherApi struct {
	Url string
}

// fetchData fetches data from the OpenWeather API
//Input: location coordinates, url
func (oapi OpenWatherApi) fetchData(location Coordinate) apiresponse {
	resp, err := http.Get(oapi.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result apiresponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	return result
}

// uses fetchData go routines to fecth the data from the API
func (oapi OpenWatherApi) fastFetchData(location Coordinate) []Weatherdata {

	c := make(chan apiresponse)
	fetchReplica := func() { c <- oapi.fetchData(location) }
	for i := 0; i < 5; i++ {
		go fetchReplica()
	}
	return (<-c).Data
}
