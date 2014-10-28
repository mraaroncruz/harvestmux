package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Response struct {
	DayEntries []Entry `json:"day_entries"`
}

type Entry struct {
	Hours float64
}

type Config struct {
	Email     string
	Password  string
	Subdomain string
}

var configPath = flag.String("config", "./config.yml", "path to config file")

const harvestUrl = "harvestapp.com/daily"

func createRequest(config *Config) *http.Request {
	url := fmt.Sprintf("https://%s.%s", config.Subdomain, harvestUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(config.Email, config.Password)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

func main() {
	flag.Parse()
	f, err := os.Open(*configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	configBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	c := Config{}
	err = yaml.Unmarshal(configBytes, &c)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	request := createRequest(&c)
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	data := Response{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	entry := data.DayEntries[0]
	fmt.Printf("%1.2f", entry.Hours)
}
