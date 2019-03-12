package main

import "log"
import "net/http"

func ping(url string, expectedCode int) (bool, error) {

	//values := url.Values{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("[WARN] Failed to initiate http request.", err)
		return false, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[WARN] Failed to get URL %v. %v", url, err)
		return false, err
	}
	defer resp.Body.Close()

	return (resp.StatusCode == expectedCode), nil
}
