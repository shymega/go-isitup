// Package isitup provides a Go wrapper around the isitup.org API
package isitup

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

/** Define JSON Response from API into struct **/
type APIResponseJSON struct {
	Domain       string  `json:"domain"`
	Port         int     `json:"port"`
	StatusCode   int     `json:"status_code"`
	ResponseIP   string  `json:"response_ip"`
	ResponseCode int     `json:"response_code"`
	ResponseTime float64 `json:"response_time"`
}

// GetStatus returns the Status Code from the API as a int.
// It requires the site and UserAgent to be specified as arguments.
func GetStatus(site, UserAgent string) int {
	url := "http://isitup.org/" + site + ".json" // Define the url to fetch with the site passed to this function

	/** Retrieve JSON response **/
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("Error creating HTTP request", "Error message: ", err)
		os.Exit(1)
	}
	req.Header.Add("User-Agent", UserAgent) // Add UserAgent passed to this function
	resp, err := client.Do(req)             // Request the URL with UserAgent passed to function
	if err != nil {
		log.Fatalln("Error retrieving JSON response from isitup.org.", "Error message: ", err)
		os.Exit(1)
	}

	/** Decode JSON Response **/
	decoder := json.NewDecoder(resp.Body)
	var data APIResponseJSON
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalln("Error decoding JSON response.", "Error message: ", err)
		os.Exit(1)
	}

	return data.StatusCode
}
