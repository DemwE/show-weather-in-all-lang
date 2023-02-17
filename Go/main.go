package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Failed to load file .env:", err)
		return
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API key is not set")
		return
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", "Warsaw", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch weather data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read data from HTTP response:", err)
		return
	}

	var data WeatherData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Failed to parse JSON data:", err)
		return
	}

	fmt.Printf("Weather in %s: %.1f°C\n", data.Name, data.Main.Temp)
}
