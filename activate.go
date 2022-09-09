package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var account string

func main() {

	videos := []string{}
	fmt.Println("Enter Video IDs: ")
	fmt.Scan(&videos)

	fmt.Println("Enter account id: ")
	fmt.Scan(&account)

	postBody := map[string]string{"state": "ACTIVE"}
	jsonValue, _ := json.Marshal(postBody)

	for _, v := range videos {

		fmt.Printf("Activating Video ID %s", v)
		req, _ := http.NewRequest("PATCH", "https://cms.api.brightcove.com/v1/accounts/"+account+"/videos/"+v, bytes.NewBuffer(jsonValue))
		req.Header.Set("Authorization", "Bearer ")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		var resp bytes.Buffer
		error := json.Indent(&resp, body, "", "\t")
		if error != nil {
			log.Println("JSON parse error:", error)
		}

		fmt.Println(resp.String())
		fmt.Print("\n")
	}
}
