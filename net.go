package mediaos

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var urlTemplate = "https://{publication}.hearst.io/api/v1/{endpoint}"

func doAPICall(endpoint Endpoint, req Request) (result []byte, err error) {
	uri := prepareAPIUri(endpoint, req)
	log.Printf("URL: %s\n", uri)

	return doGet(uri)
}

func prepareAPIUri(endpoint Endpoint, req Request) (uri string){
	uri = strings.Replace(urlTemplate, "{publication}", string(req.publication), 1)
	uri = strings.Replace(uri, "{endpoint}", string(endpoint), 1)

	params := prepareParams(req.key, req)
	uri += "?" + params.Encode()

	return
}

func doGet(url string) (result []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, Error(resp.StatusCode)
	}

	return result, nil
}
