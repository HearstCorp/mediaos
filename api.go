package mediaos

import (
	"encoding/json"
)

var MediaOs = &mediaOsAPI{}

type Client interface {
	GetContent(publication PubData, key string, endpoint Endpoint, req Request) (res ContentResponse, err error)
	GetImages(publication PubData, key string, req Request) (res ImageResponse, err error)
}

type mediaOsAPI struct{}

func (m *mediaOsAPI) GetContent(publication PubData, key string, endpoint Endpoint, req Request) (res ContentResponse, err error) {
	addRequestContext(publication, key, &req)

	bytes, err := doAPICall(endpoint, req)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (m *mediaOsAPI) GetImages(publication PubData, key string, req Request) (res ImageResponse, err error) {
	addRequestContext(publication, key, &req)

	bytes, err := doAPICall(ImagesAPI, req)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func addRequestContext(publication PubData, key string, req *Request) {
	req.key = key
	req.publication = publication
}
