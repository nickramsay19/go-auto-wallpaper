package main

import (
	"fmt" // for debug printing
	"encoding/json" // for reading "config.json"
	"io"
	"io/ioutil"
	"os"
	"log"
	"net/http"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(url string, filepath string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {

	// read the local "config.json" file to retrieve the API key
	configBody, err := ioutil.ReadFile("./config.json")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

	// parse/unmarshal the json bytes to a generic json object
	var json_object map[string]interface{} // define a generic json object that maps json values to its keys
	if err := json.Unmarshal(configBody, &json_object); err != nil {
		log.Fatalln(err)
	}
	apiKey := json_object["api_key"].(string)

	// create our request from the apiKey
	req := NewRequest("https://api.unsplash.com/photos/random")
	//req.SetParameter("collections", "1")
	req.SetParameter("orientations", "landscape")
	req.SetParameter("client_id", apiKey)

	resBytes, err := req.GetResponse()
	if err != nil {
		log.Fatalln(err)
	}

	// convert the resBody to json
	var resJson map[string]interface{} // generic object to hold json values
	json.Unmarshal(resBytes, &resJson)

	// extract response data
	imageId := resJson["id"].(string) // get the id for naming
	// to get the url we must first get the "urls" object
	urlsJson := resJson["urls"].(map[string]interface{}) // access the json object "urls" which contains the url strings
	imageUrlRaw := urlsJson["raw"].(string) // access the "raw" key in the "urls" object

	// download the file
	imageFilename := fmt.Sprintf("./%s.png", imageId)
	DownloadFile(imageUrlRaw, imageFilename)
}
