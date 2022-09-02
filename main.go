package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gmvivekanandan/go-weather-cli/weather"
)

var place string
var country string

func init() {
	flag.StringVar(&place, "place", "", "place to display weather")
	flag.StringVar(&country, "country", "", "country the place belongs to")
}

func main() {

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: weathergo -place <place> -country <country>\n")
	}

	if len(place) == 0 {
		flag.Usage()
		os.Exit(1)
	} else if len(country) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	lat, lon := weather.GetCoordinates(place, country)
	currentWeather := weather.GetCurrentWeather(lat, lon)
	placeTime := (time.Now()).Local()

	fmt.Println("Current date and time is: ", placeTime.String())
	fmt.Println("Temperature:", currentWeather.Main.Temp, "°C")
}
