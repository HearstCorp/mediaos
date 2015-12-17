package mediaos

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetApiPath(publication PubData, endpoint Endpoint, params map[string]string) (uri string) {
	if nil == publication {
		return
	}

	uri = strings.Replace(Config.urlTemplate, "{protocol}", Config.protocol, 1)
	uri = strings.Replace(uri, "{domainPort}", publication.MosDomainAndPort(), 1)
	uri = strings.Replace(uri, "{endpoint}", endpoint.String(), 1)

	p := url.Values{}
	for key, value := range params {
		p.Set(key, value)
	}

	uri += "?" + p.Encode()

	return
}

func doAPICall(endpoint Endpoint, req Request) (result []byte, err error) {
	uri := prepareAPIUri(endpoint, req)
	log.Printf("URL: %s\n", uri)

	return doGet(uri)
}

func prepareAPIUri(endpoint Endpoint, req Request) (uri string) {
	if nil == req.publication {
		return
	}

	uri = strings.Replace(Config.urlTemplate, "{protocol}", Config.protocol, 1)
	uri = strings.Replace(uri, "{domainPort}", req.publication.MosDomainAndPort(), 1)
	uri = strings.Replace(uri, "{endpoint}", endpoint.String(), 1)

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
