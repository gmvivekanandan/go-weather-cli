package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type CoordinatesType struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

type CurrentWeatherType struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func GetCoordinates(place string, country string) (float64, float64) {

	var CoordinateLocation []CoordinatesType
	var Lat float64
	var Lon float64

	response, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=" + place + "," + country + "&limit=1&appid=" + "89b7974c56e87de5ba3f940b66a2a339")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &CoordinateLocation)

	for _, val := range CoordinateLocation {
		Lat = val.Lat
		Lon = val.Lon
	}

	return Lat, Lon
}

func GetCurrentWeather(latitude float64, longitude float64) CurrentWeatherType {

	var CurrentWeather CurrentWeatherType

	lat := fmt.Sprintf("%f", latitude)
	lon := fmt.Sprintf("%f", longitude)
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + string(lat) + "&lon=" + lon + "&appid=" + "89b7974c56e87de5ba3f940b66a2a339&units=metric")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &CurrentWeather)
	return CurrentWeather
}
