package mediaos

// Client is the API client interface
type Client interface {
	Get(Endpoint, Request) (Response, error)
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

func (c *client) Get(endpoint Endpoint, req Request) (Response, error) {
	c.addRequestContext(&req)
	return doAPICall(endpoint, req)
}

func (c *client) addRequestContext(req *Request) {
	req.key = c.key
	req.publication = c.publication
}
