package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type weather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

type pass struct {
	Success bool
	weather
}

func getData(city string) weather {
	key := "http://api.weatherapi.com/v1/current.json?key=6c62ced4a680400daeb104136211901&q="
	url := key + city
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var all weather
	err = json.Unmarshal([]byte(body), &all)
	if err != nil {
		fmt.Println(err)
	}

	file, err := json.MarshalIndent(all, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("data.json", file, 064)
	if err != nil {
		fmt.Println(err)
	}
	return all
}

func display(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, pass{true, getData(r.FormValue("city"))})
	} else {
		tmpl.Execute(w, pass{false, getData(r.FormValue("city"))})
	}
}

func main() {
	http.HandleFunc("/", display)
	http.ListenAndServe(":8080", nil)
}
