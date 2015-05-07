package mediaos

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var urlTemplate = "https://{publication}.hearst.io/api/v1/{endpoint}"

func doAPICall(endpoint Endpoint, req Request) (res Response, err error) {

	uri := strings.Replace(urlTemplate, "{publication}", string(req.publication), 1)
	uri = strings.Replace(uri, "{endpoint}", string(endpoint), 1)

	params := prepareParams(req.key, req)
	uri += "?" + params.Encode()

	log.Printf("URL: %s\n", uri)

	bytes, err := doGet(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}

	return res, nil
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
