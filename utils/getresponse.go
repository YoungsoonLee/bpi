package utils

import (
	"io/ioutil"
	"net/http"
)

// GetResponse ...
// Call url and read response body and return responsebody []byte
func GetResponse(url string) ([]byte, error) {
	// call url
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// read response body
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
