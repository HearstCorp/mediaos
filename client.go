package mediaos

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

func (c *client) GetContent(req Request) (Response, error) {
	c.addRequestContext(&req)
	return doAPICall(ContentAPI, req)
}

func (c *client) addRequestContext(req *Request) {
	req.key = c.key
	req.publication = c.publication
}
