package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func hello(response http.ResponseWriter, _ *http.Request) {
	_, err := response.Write([]byte("Hi from Golang..."))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(response http.ResponseWriter, request *http.Request) {
		// get the name of the city
		cityName := strings.SplitN(request.URL.Path, "/", 3)[2]
		// make the API call
		cityInfo, err := getCityWeatherInfo(cityName)
		// validate the data
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Header().Set("Content-type", "applications/json; charset=utf-8")
		err = json.NewEncoder(response).Encode(cityInfo)
		if err != nil {
			return
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func getCityWeatherInfo(cityName string) (weatherData, error) {
	apiConfig, err := loadApiConfig("./.apiConfig")
	if err != nil {
		return weatherData{}, err
	}
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + cityName)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if err != nil {
		return weatherData{}, err
	}

	var data weatherData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}
	return data, nil
}
