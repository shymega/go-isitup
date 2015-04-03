// Package isitup provides a Go wrapper around the isitup.org API
package isitup

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// APIResponseJSON is a struct defining the JSON response returned
// from isitup.org
type APIResponseJSON struct {
	Domain       string  `json:"domain"`
	Port         int     `json:"port"`
	StatusCode   int     `json:"status_code"`
	ResponseIP   string  `json:"response_ip"`
	ResponseCode int     `json:"response_code"`
	ResponseTime float64 `json:"response_time"`
}

var data APIResponseJSON

func getJSONresp(site, UserAgent string) {
	url := "http://isitup.org/" + site + ".json"
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

	// Decode JSON response
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalln("Error decoding JSON response.", "Error message: ", err)
		os.Exit(1)
	}
}

// GetDomain returns the domain requested, requires site and UserAgent as arguments (both strings).
// This function is probably unnecessary, seeing as the domain is already passed to the function itself.
func GetDomain(site, UserAgent string) string {
	getJSONresp(site, UserAgent)
	return data.Domain
}

// GetIP returns the IP of the site, requires site and UserAgent as arguments (both strings).
func GetIP(site, UserAgent string) string {
	getJSONresp(site, UserAgent)
	return data.ResponseIP
}

// GetPort returns the port of the site (usually 80 for HTTP), requires site and UserAgent as arguments (both strings).
func GetPort(site, UserAgent string) int {
	getJSONresp(site, UserAgent)
	return data.Port
}

// GetStatus returns the HTTP status code from the website, requires site and UserAgent as arguments (both strings)
func GetStatus(site, UserAgent string) int {
	getJSONresp(site, UserAgent)
	return data.StatusCode
}

// GetResponseTime returns the time taken for a response from the site specifed, requires site and UserAgent as arguments (both strings)
func GetResponseTime(site, UserAgent string) float64 {
	getJSONresp(site, UserAgent)
	return data.ResponseTime
}
