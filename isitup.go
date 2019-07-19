// Package isitup provides a Go wrapper around the isitup.org API

/*
*  Copyright 2018 Dom Rodriguez (shymega)
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
*/

package isitup // import "github.com/shymega/go-isitup"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// APIResponse is a model of the data returned
// from isitup.org
type APIResponse struct {
	Domain       string  `json:"domain"`
	Port         int     `json:"port"`
	StatusCode   int     `json:"status_code"`
	ResponseIP   string  `json:"response_ip"`
	ResponseCode int     `json:"response_code"`
	ResponseTime float64 `json:"response_time"`
}

// getAPIResponse returns a JSON struct of the API's response.
// It requires a site to test upon, and a UserAgent.
// This function is used internally.
func getAPIResponse(site, UserAgent string) (response APIResponse) {
	// Instantiate a new instance of the JSON data model.
	data := APIResponse{}

	// Format the request URL
	url := fmt.Sprintf("http://isitup.org/%s.json", site)

	// Instantiate a new client
	client := &http.Client{}

	// Retrieve the API response
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("Error creating HTTP request", "Error message: ", err)
		os.Exit(1)
	}

	// Add the UserAgent to the HTTP request headers
	req.Header.Add("User-Agent", UserAgent)

	// Request the URL
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error retrieving JSON response", "Error message: ", err)
		os.Exit(1)
	}

	// Decode JSON response into the `data` instance.
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalln("Error decoding JSON response", "Error message: ", err)
		os.Exit(1)
	}

	// Return the `data` model from the API.
	return data
}

// GetDomain returns the domain requested, requires site and UserAgent
// as arguments (both strings). This function is probably
// unnecessary, seeing as the domain is already passed to the function
// itself.
func GetDomain(site, UserAgent string) string {
	return getAPIResponse(site, UserAgent).Domain
}

// GetIP returns the IP of the site, requires site and UserAgent as
// arguments (both strings).
func GetIP(site, UserAgent string) string {
	return getAPIResponse(site, UserAgent).ResponseIP
}

// GetPort returns the port of the site (usually 80 for HTTP),
// requires site and UserAgent as arguments (both strings).
func GetPort(site, UserAgent string) int {
	return getAPIResponse(site, UserAgent).Port
}

// GetStatusCode returns the HTTP status code from the website, requires
// site and UserAgent as arguments (both strings)
func GetStatusCode(site, UserAgent string) int {
	return getAPIResponse(site, UserAgent).ResponseCode
}

// GetResponseTime returns the time taken for a response from the site
// specified, requires site and UserAgent as arguments (both strings)
func GetResponseTime(site, UserAgent string) float64 {
	return getAPIResponse(site, UserAgent).ResponseTime
}
