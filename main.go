package main

import (
	//"flag"
	"encoding/json"
	"fmt"
	"github.com/reujab/wallpaper"
	"io/ioutil"
	"net/http"
)

const apiKey = "rnmuXya1IwVfoHiXiAsWkLDbIfDHVN1GsvT1QuVg"

type APODResponse struct {
	HDURL string `json:"hdurl"`
}

func getImage(nasaApiKey string) (string, error) {
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s", nasaApiKey)

	// Send GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(body), nil
}

func main() {
	var response APODResponse
	receivedJson, err := getImage(apiKey)

	if err != nil {
		fmt.Errorf("error when fetching image: %v", err)
		return
	}

	err = json.Unmarshal([]byte(receivedJson), &response)
	if err != nil {
		fmt.Errorf("Error: %v", err)
		return
	}

	err = wallpaper.SetFromURL(response.HDURL)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return
	}

	err = wallpaper.SetMode(wallpaper.Center)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return
	}

	fmt.Println("Enjoy your new spaceouse wallpaper!")
}
