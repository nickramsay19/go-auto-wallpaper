package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestParameter struct {
	key string
	value string
}

type Request struct {
	url string
	params []RequestParameter
}

func NewRequest(url string) (req Request) {

	// initialse a params slice
	params := make([]RequestParameter, 0)

	// create the request
	req = Request{url, params}

	return req
}

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

func (req Request) GetResponse() (resBody []byte, err error) {
	
	// construct the request url from url and params
	url := req.url
	fmt.Println(len(req.params))
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