package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	api_token := os.Getenv("EVENTS_API_TOKEN")
	url := "https://events.1password.com"

	start_time := time.Now().AddDate(0, 0, -1)

	payload := []byte(fmt.Sprintf(`{
		"limit": 20,
		"start_time": "%s"
	}`, start_time.Format(time.RFC3339)))

	client := &http.Client{}

	signinsRequest, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/signinattempts", url), bytes.NewBuffer(payload))
	signinsRequest.Header.Set("Content-Type", "application/json")
	signinsRequest.Header.Set("Authorization", "Bearer "+api_token)
	signinsResponse, signinsError := client.Do(signinsRequest)
	if signinsError != nil {
		panic(signinsError)
	}
	defer signinsResponse.Body.Close()
	signinsBody, _ := ioutil.ReadAll(signinsResponse.Body)
	fmt.Println(string(signinsBody))

	usagesRequest, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/itemusages", url), bytes.NewBuffer(payload))
	usagesRequest.Header.Set("Content-Type", "application/json")
	usagesRequest.Header.Set("Authorization", "Bearer "+api_token)
	usagesResponse, usagesError := client.Do(usagesRequest)
	if usagesError != nil {
		panic(usagesError)
	}
	defer usagesResponse.Body.Close()
	usagesBody, _ := ioutil.ReadAll(usagesResponse.Body)
	fmt.Println(string(usagesBody))
}
