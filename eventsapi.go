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
	api_token := "Bearer " + os.Getenv("EVENTS_API_TOKEN")
	url := "https://events.1password.com"

	start_time := time.Now().AddDate(0, 0, -1)

	payload := []byte(fmt.Sprintf(`{
		"limit": 20,
		"start_time": "%s"
	}`, start_time.Format(time.RFC3339)))

	request, error := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/signinattempts", url), bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", api_token)
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
