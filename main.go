package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type weather struct {
	Check    bool
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
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func display(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := "http://api.weatherapi.com/v1/current.json?key=6c62ced4a680400daeb104136211901&q="
	city := r.FormValue("city")
	url := key + city

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		tmpl.Execute(w, nil)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		tmpl.Execute(w, nil)
	}
	var data weather
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
		tmpl.Execute(w, nil)
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
		tmpl.Execute(w, nil)
	}
	err = ioutil.WriteFile("data.json", file, 064)
	if err != nil {
		fmt.Println(err)
		tmpl.Execute(w, nil)
	}

	if data.Error.Code == 1006 {
		data.Check = false
		tmpl.Execute(w, data)
	} else {
		data.Check = true
		tmpl.Execute(w, data)
	}

}

func main() {
	fmt.Println("Server - http://localhost:8080/")
	http.HandleFunc("/", display)
	http.ListenAndServe(":8080", nil)
}
