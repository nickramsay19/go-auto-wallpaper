package main

import (
	"fmt" // for debug printing
	"encoding/json" // for reading "config.json"
	"io/ioutil"
	"log"
	"os" // for command line args
)

func main() {

	// get the command line arguments
	args := os.Args[1:] // scrap the program name
	var fileNameSpecified bool = len(args) > 0 // check if enough args are specified for a filename

	// read the local "config.json" file to retrieve the API key
	configBody, err := ioutil.ReadFile("/Users/nickramsay/Documents/Projects/go-auto-wallpaper/secret.json")
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
	req.SetParameter("topics", "bo8jQKTaE0Y")
	req.SetParameter("orientation", "landscape")
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

	// create the filename of the file to be downloaded
	var imageFilename string
	if fileNameSpecified {
		imageFilename = fmt.Sprintf("./%s", args[0]) // use the name passed from cl arguments
	} else {
		imageFilename = fmt.Sprintf("./%s.png", imageId) // use the id retreived from the api
	}
	
	// download the file
	DownloadFile(imageUrlRaw, imageFilename)
}
