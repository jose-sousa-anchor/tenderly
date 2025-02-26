package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Vnet struct {
	Id string `json:"id"`
}

const TenderlyOptimismURL = "https://api.tenderly.co/api/v1/account/anchorage/project/optimism-regnet/vnets"
const TenderlyBaseURL = "https://api.tenderly.co/api/v1/account/anchorage/project/base-regnet/vnets"
const TenderlyBobURL = "https://api.tenderly.co/api/v1/account/anchorage/project/bob-regnet/vnets"

var tenderlyProjectsUrls = []string{TenderlyOptimismURL, TenderlyBaseURL, TenderlyBobURL}

func main() {
	client := &http.Client{}

	for _, tenderlyProjectUrl := range tenderlyProjectsUrls {
		// Create GET request
		req, err := http.NewRequest("GET", tenderlyProjectUrl, nil)
		if err != nil {
			fmt.Println("Error creating GET request:", err)
			return
		}

		req.Header.Add("X-Access-Key", "EharCQ2AJxJoFgbIL3Yx3kcMNcJlSQTE")

		// Add query parameters
		q := req.URL.Query()
		q.Add("perPage", "20")
		q.Add("page", "2")
		req.URL.RawQuery = q.Encode()

		fmt.Println("Request URL:", req.URL.String())

		// Execute GET request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error executing GET request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error: GET request failed with status", resp.StatusCode)
			return
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Parse JSON response
		var vnets []Vnet
		if err := json.Unmarshal(body, &vnets); err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}

		// Loop through Vnets and send DELETE requests
		for _, v := range vnets {
			uri := fmt.Sprintf("%s/%s", tenderlyProjectUrl, v.Id)

			req, err := http.NewRequest("DELETE", uri, nil)
			if err != nil {
				fmt.Println("Error creating DELETE request:", err)
				continue
			}

			req.Header.Add("X-Access-Key", "EharCQ2AJxJoFgbIL3Yx3kcMNcJlSQTE")

			// Execute DELETE request
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error executing DELETE request:", err)
				continue
			}
			defer resp.Body.Close()

			fmt.Printf("Deleted Vnet %s - Status: %d\n", v.Id, resp.StatusCode)
		}
	}
}
