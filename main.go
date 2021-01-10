package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PetrNovak96/gohttpclient/model"
	"io/ioutil"
	"net/http"
	"os"
)

func LoadConfiguration(filename string) (model.Config, error) {
	var config model.Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config);
	return config, err
}

func main() {

	// TODO potřebuju tam mít validaci ... playground validator package
	fmt.Println("Starting the application...")
	config, _ := LoadConfiguration("config.json")
	fmt.Println(config)

	// GET
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// POST
	jsonData := map[string]string{"firstname": "Petr", "lastname": "Novák"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// POST jinak
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	i := model.Item{}
	i.Shout()
}