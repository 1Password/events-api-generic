package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

// For more information, check out our support page
// https://support.1password.com/events-reporting

// Replace APITOKEN with your generated Events API token
const token = "Bearer APITOKEN"

// Replace url with the URL corresponding to your 1Password account region
const url = "https://events.1password.com"

// For more information on the response, check out our support page
// https://support.1password.com/cs/events-api-reference/
type SignInAttemptResponse struct {
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"has_more"`
	Items   []struct {
		UUID        string    `json:"uuid"`
		SessionUUID string    `json:"session_uuid"`
		Timestamp   time.Time `json:"timestamp"`
		Country     string    `json:"country"`
		Category    string    `json:"category"`
		Type        string    `json:"type"`
		Details     *struct {
			Value string `json:"value"`
		} `json:"details"`
		SignInAttemptTargetUser struct {
			UUID  string `json:"uuid"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"target_user"`
		SignInAttemptClient struct {
			AppName         string `json:"app_name"`
			AppVersion      string `json:"app_version"`
			PlatformName    string `json:"platform_name"`
			PlatformVersion string `json:"platform_version"`
			OSName          string `json:"os_name"`
			OSVersion       string `json:"os_version"`
			IPAddress       string `json:"ip_address"`
		} `json:"client"`
	} `json:"items"`
}

type ItemUsageResponse struct {
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"has_more"`
	Items   []struct {
		UUID          string    `json:"uuid"`
		Timestamp     time.Time `json:"timestamp"`
		UsedVersion   uint32    `json:"used_version"`
		VaultUUID     string    `json:"vault_uuid"`
		ItemUUID      string    `json:"item_uuid"`
		ItemUsageUser struct {
			UUID  string `json:"uuid"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"user"`
		ItemUsageClient struct {
			AppName         string `json:"app_name"`
			AppVersion      string `json:"app_version"`
			PlatformName    string `json:"platform_name"`
			PlatformVersion string `json:"platform_version"`
			OSName          string `json:"os_name"`
			OSVersion       string `json:"os_version"`
			IPAddress       string `json:"ip_address"`
		} `json:"client"`
	} `json:"items"`
}

func SignInAttemptWorker(stop chan os.Signal, ctx context.Context) error {
	httpClient := &http.Client{}
	cursor := ""
	pollingTime := 5 * time.Second
	ticker := time.NewTicker(pollingTime)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-stop:
			return nil
		case <-ticker.C:
			for {
				var body io.Reader
				if cursor == "" {
					// Start loading events from 24 hours ago with a limit of 20 per response
					now := time.Now()
					now = now.Add(-24 * time.Hour)
					body = strings.NewReader(fmt.Sprintf("{\"limit\": 20, \"start_time\": \"%s\"}", now.Format(time.RFC3339)))
				} else {
					// Use the existing cursor to poll any new events
					body = strings.NewReader(cursor)
				}
				request, _ := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", url, "/api/v1/signinattempts"), body)
				request.Header.Add("Authorization", token)
				request.Header.Add("User-Agent", "1Password Events API Script")
				response, err := httpClient.Do(request)
				if err != nil {
					return err
				}
				if response.StatusCode != 200 {
					return fmt.Errorf("Unexpected Status Code: %s", response.Status)
				}

				var signInAttemptsResponse SignInAttemptResponse
				err = json.NewDecoder(response.Body).Decode(&signInAttemptsResponse)
				if err != nil {
					return fmt.Errorf("Failed to unmarshal response: %w", err)
				}

				cursor = fmt.Sprintf(`{ "cursor": "%s" }`, signInAttemptsResponse.Cursor)

				for _, item := range signInAttemptsResponse.Items {
					i, _ := json.Marshal(item)
					fmt.Println(string(i))
				}

				if !signInAttemptsResponse.HasMore {
					break
				}
			}
		}
	}
}

func ItemUsageWorker(stop chan os.Signal, ctx context.Context) error {
	httpClient := &http.Client{}
	cursor := ""
	pollingTime := 5 * time.Second
	ticker := time.NewTicker(pollingTime)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-stop:
			return nil
		case <-ticker.C:
			for {
				var body io.Reader
				if cursor == "" {
					// Start loading events from 24 hours ago with a limit of 20 per response
					now := time.Now()
					now = now.Add(-24 * time.Hour)
					body = strings.NewReader(fmt.Sprintf("{\"limit\": 20, \"start_time\": \"%s\"}", now.Format(time.RFC3339)))
				} else {
					// Use the existing cursor to poll any new events
					body = strings.NewReader(cursor)
				}
				request, _ := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", url, "/api/v1/itemusages"), body)
				request.Header.Add("Authorization", token)
				request.Header.Add("User-Agent", "1Password Events API Script")
				response, err := httpClient.Do(request)
				if err != nil {
					return err
				}
				if response.StatusCode != 200 {
					return fmt.Errorf("Unexpected Status Code: %s", response.Status)
				}

				var itemUsageResponse ItemUsageResponse
				err = json.NewDecoder(response.Body).Decode(&itemUsageResponse)
				if err != nil {
					return fmt.Errorf("Failed to unmarshal response: %w", err)
				}

				cursor = fmt.Sprintf(`{ "cursor": "%s" }`, itemUsageResponse.Cursor)

				for _, item := range itemUsageResponse.Items {
					i, _ := json.Marshal(item)
					fmt.Println(string(i))
				}

				if !itemUsageResponse.HasMore {
					break
				}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	sigsChan := make(chan os.Signal)
	signal.Notify(sigsChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := SignInAttemptWorker(sigsChan, ctx)
		if err != nil {
			fmt.Println(err)
		}
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := ItemUsageWorker(sigsChan, ctx)
		if err != nil {
			fmt.Println(err)
		}
		cancel()
	}()

	wg.Wait()
}
