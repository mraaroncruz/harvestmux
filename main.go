package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type response struct {
	DayEntries []Entry `json:"day_entries"`
}

type Entry struct {
	Client            string      `json:"client"`
	CreatedAt         time.Time   `json:"created_at"`
	Hours             float64     `json:"hours"`
	HoursWithoutTimer float64     `json:"hours_without_timer"`
	ID                int         `json:"id"`
	Notes             interface{} `json:"notes"`
	Project           string      `json:"project"`
	ProjectID         string      `json:"project_id"`
	SpentAt           string      `json:"spent_at"`
	Task              string      `json:"task"`
	TimerStartedAt    *time.Time  `json:"timer_started_at"`
	TaskID            string      `json:"task_id"`
	UpdatedAt         time.Time   `json:"updated_at"`
	UserID            int         `json:"user_id"`
}

type config struct {
	Email     string
	Password  string
	Subdomain string
}

var configPath = flag.String("config", "./config.yml", "path to config file")
var currentOnly = flag.Bool("o", false, "only show current time")

const harvestURL = "harvestapp.com/daily?slim=1"

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
	for _, e := range data.DayEntries {
		sum += e.Hours
	}

	for _, e := range data.DayEntries {
		if e.TimerStartedAt != nil {
			fmt.Printf("%s — %1.2f - %1.2f", e.Client, e.Hours, sum)
			return
		}
	}
	fmt.Printf("No Timer Running!")
}
