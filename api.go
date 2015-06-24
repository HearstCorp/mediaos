package mediaos

import (
	"encoding/json"
)

type Client interface {
	GetArticles(Endpoint, Request) (ArticleResponse, error)
	GetImages(Endpoint, Request) (ImageResponse, error)
}

// New creates a new API client object for the given publication
func New(pub Publication, key string) Client {
	return &client{
		publication: pub,
		key:         key,
	}
}

type client struct {
	publication Publication
	key         string
}

func (c *client) get(endpoint Endpoint, req Request) (result []byte, err error) {
	c.addRequestContext(&req)
	return doAPICall(endpoint, req)
}

func (c *client) GetArticles(endpoint Endpoint, req Request) (res ArticleResponse, err error) {
	bytes, err := c.get(endpoint, req)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *client) GetImages(endpoint Endpoint, req Request) (res ImageResponse, err error) {
	bytes, err := c.get(endpoint, req)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *client) addRequestContext(req *Request) {
	req.key = c.key
	req.publication = c.publication
}
