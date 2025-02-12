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

func main() {
	client := &http.Client{}
	var page int = 10
	for page > 1 {
		req, _ := http.NewRequest("GET", "https://api.tenderly.co/api/v1/account/anchorage/project/optimism-regnet/vnets", nil)
		req.Header.Add("X-Access-Key", "EharCQ2AJxJoFgbIL3Yx3kcMNcJlSQTE")
		q := req.URL.Query()
		q.Add("perPage", "20")
		req.URL.RawQuery = q.Encode()
		fmt.Printf("getting page %d\n", page)
		resp, _ := client.Do(req)
		fmt.Println(resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		vnets := []Vnet{}
		err := json.Unmarshal(body, &vnets)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v := range vnets {
			uri := fmt.Sprintf("https://api.tenderly.co/api/v1/account/anchorage/project/optimism-regnet/vnets/%s", v.Id)
			req, _ = http.NewRequest("DELETE", uri, nil)
			req.Header.Add("X-Access-Key", "EharCQ2AJxJoFgbIL3Yx3kcMNcJlSQTE")
			resp, _ := client.Do(req)
			fmt.Println(resp.StatusCode)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		page--
	}
}
