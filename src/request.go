package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// a simple keyvalue struct for storing each request parameter
type RequestParameter struct {
	key string
	value string
}

// store the request as a struct holding its url and params separately
type Request struct {
	url string
	params []RequestParameter
}

// request constructor/initialiser
func NewRequest(url string) (req Request) {

	// initialse a params slice
	params := make([]RequestParameter, 0)

	// create the request
	req = Request{url, params}

	return req
}

// add or set a request paramter
func (req *Request) SetParameter(key string, value string) {

	// first check if the parameter already exists
	for _, s := range req.params {
		if (key == s.key) {
			s.value = value // change the value
			return // we set the value we can finish now
		}
	}

	// we didn't find it, add it
	// create the new RequestParamater
	newParam := RequestParameter{key, value}

	// append to the req
	req.params = append(req.params, newParam)

	return
}

// fetch the api response from this request
func (req Request) GetResponse() (resBody []byte, err error) {
	
	// construct the request url from url and params
	url := req.url
	for i, p := range req.params {

		// parameters are separated by "&" or "?" on the first param
		if i == 0 {
			url += "?"
		} else {
			url += "&"
		}

		// add the key and value
		url += p.key + "=" + p.value
	}

	// send the request
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// read the body
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
	   log.Fatalln(err)
	   return nil, err
	}

	// no error
	return resBody, nil
}