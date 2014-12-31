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

type response struct {
	DayEntries []entry `json:"day_entries"`
}

type entry struct {
	Hours float64
}

type config struct {
	Email     string
	Password  string
	Subdomain string
}

var configPath = flag.String("config", "./config.yml", "path to config file")
var currentOnly = flag.Bool("o", false, "only show current time")

const harvestURL = "harvestapp.com/daily"

func createRequest(c *config) *http.Request {
	url := fmt.Sprintf("https://%s.%s", c.Subdomain, harvestURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(c.Email, c.Password)
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

	c := config{}
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

	data := response{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	var sum float64
	if *currentOnly {
		entry := data.DayEntries[len(data.DayEntries)-1]
		sum = entry.Hours
	} else {
		for _, entry := range data.DayEntries {
			sum += entry.Hours
		}
	}
	fmt.Printf("%1.2f", sum)
}
