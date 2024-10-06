package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	url := "http://localhost:8080"
	response, err := TryConnection(url)
	if err != nil {
		fmt.Println("Error while trying to connect. Please try again!") // This is not the message to show the user. Change to a user facing message later.
		return
	}
	// At this point, the connection has succeeded.
	// Try connections until StatusCode is 200 (OK)
	// fmt.Println("response StatusCode: ", response.StatusCode)
	for response.StatusCode != 200 {
		if response.StatusCode == 429 { // Recieved a Retry-After header.
			r := response.Header.Get("Retry-After") // Could be buggy and be a string, but should be an int
			ir, err := strconv.Atoi(r)              // Try to convert to an int
			if err != nil {                         // If we recieve a buggy response from the Retry-After header, just set to 1 second
				ir = 1
			}
			if ir <= 5 { // If the wait time is less than 5 seconds, retry
				if ir > 1 {
					fmt.Println("Retrying connection, this may take a second.")
				}
				response, err = TryConnection(url)
				if err != nil {
					fmt.Println("Error while trying to connect. Please try again!")
					return
				}
			} else { // If the wait time is longer than 5 seconds, just exit
				fmt.Println("Quitting because connection is taking too long. Please try again!")
				return
			}
		}
	}
	//	fmt.Println("response Status: ", response.Status)
	//	fmt.Println("response StatusCode: ", response.StatusCode)

	// Read response body (it is a byte array)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	// Convert byte array to string and print it
	responseString := string(body)
	fmt.Println("Weather today: ", responseString)

}
func TryConnection(url string) (resp *http.Response, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
