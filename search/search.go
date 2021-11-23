package search

import (
	"time"
)

type Conditions struct {
	Temp     float32
	Humidity int
}

func search(conditions Conditions, data []Weatherdata) []Weatherdata {
	var result []Weatherdata = make([]Weatherdata, 0, len(data))
	for _, val := range data {
		if val.Main.Temp <= conditions.Temp && val.Main.Humidity >= conditions.Humidity {
			result = append(result, val)
		}
	}
	return result
}

//launch search go routines for slices of the data
func (oapi OpenWatherApi) QuickSearch(location Coordinate, conditions Conditions) []Weatherdata {
	data := oapi.fastFetchData(location)
	c := make(chan []Weatherdata)
	var result []Weatherdata = make([]Weatherdata, 0, len(data))
	for i := 0; i < len(data)/5; i++ {
		go func() {
			c <- search(conditions, data[:5])
		}()
	}
	timeout := time.After(1000 * time.Millisecond)
	for i := 0; i < len(data)/5; i++ {
		select {
		case res := <-c:
			result = append(result, res...)
		case <-timeout:
			panic("Timeout")
		}
	}
	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))
	return result
}
