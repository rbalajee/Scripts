package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var apikey string

type Result struct {
	Job []Jobs `json:"jobs"`
}

type Jobs struct {
	ID string `json:"id"`
}

type State struct {
	Job struct {
		StateHistory []StateHistory `json:"state_history"`
	} `json:"job"`
}

type StateHistory struct {
	State               string `json:"state"`
	TransitionTimestamp int64  `json:"transition_timestamp"`
}

func main() {

	fmt.Print("Enter API KEY: ")
	fmt.Scan(&apikey)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.bcovlive.io/v1/jobs?page_size=1000", nil)
	req.Header.Add("X-API-KEY", apikey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var out Result
	err = json.Unmarshal(body, &out)
	if err != nil {
		log.Println("JSON MARSHAL ERROR: ", err)
	}

	for _, v := range out.Job {
		fmt.Println("\nJob ID:", v.ID)
		req, _ := http.NewRequest("GET", "https://api.bcovlive.io/v1/jobs/"+v.ID, nil)
		req.Header.Set("X-API-KEY", apikey)

		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		var output State
		json.Unmarshal(body, &output)

		for _, val := range output.Job.StateHistory {
			fmt.Println()
			fmt.Println(val.State, "\t", time.UnixMilli(val.TransitionTimestamp))
		}

		fmt.Print("\n")
	}
}
