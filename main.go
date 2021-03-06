package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

var (
	Key string
)

type Data struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	ICON        string `json:"icon"`
}

type Main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

func init() {
	flag.StringVar(&Key, "key", "", "Api key")
	flag.Parse()
}

func main() {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=nagano&appid=%s", Key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error get api:", err)
		return
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for {
		var data Data
		err := decoder.Decode(&data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decode json:", err)
			return
		}
		fmt.Println(data)
	}
}
