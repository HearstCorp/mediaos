package mediaos

import (
	"encoding/json"
)
/*
type Client interface {
	GetContent(Endpoint, Request) (ContentResponse, error)
	GetImages(Request) (ImageResponse, error)
}

// New creates a new API client object for the given publication
func New(pub PubData, key string) Client {
	return &client{
		publication: pub,
		key:         key,
	}
}

type client struct {
	publication PubData
	key         string
}
*/
/*
func (c *client) get(endpoint Endpoint, req Request) (result []byte, err error) {
	c.addRequestContext(&req)
	return doAPICall(endpoint, req)
}
*/
func /*(c *client)*/ GetContent(publication PubData, key string, endpoint Endpoint, req Request) (res ContentResponse, err error) {
	addRequestContext(publication, key, req)

	bytes, err := /*c.*/doAPICall(endpoint, req)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func /*(c *client)*/ GetImages(publication PubData, key string, req Request) (res ImageResponse, err error) {
	addRequestContext(publication, key, req)

	bytes, err := /*c.*/doAPICall(ImagesAPI, req)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func /*(c *client)*/ addRequestContext(publication PubData, key string, req Request) {
	req.key = key
	req.publication = publication
}
