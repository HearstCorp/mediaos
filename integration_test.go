// +build integration

package mediaos

import (
	"flag"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var key = flag.String("apikey", "", "API key to use for integration testing")

func Test_GetContent(t *testing.T) {
	client := New(Cosmo, *key)

	time.Sleep(time.Second)
	response, err := client.Get(ContentAPI, Request{
		Title: "sex",
	})
	assert.Nil(t, err)
	assert.Equal(t, DefaultLimit, response.Count)
	assert.Equal(t, DefaultLimit, len(response.Items))
}

func Test_GetArticles(t *testing.T) {
	client := New(Cosmo, *key)

	time.Sleep(time.Second)
	response, err := client.Get(ArticlesAPI, Request{
		Title: "sex",
	})
	assert.Nil(t, err)
	assert.Equal(t, DefaultLimit, response.Count)
	assert.Equal(t, DefaultLimit, len(response.Items))

	// get by ID
	id := response.Items[0].ID
	time.Sleep(time.Second)
	response, err = client.Get(ArticlesAPI, Request{
		ID: id,
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, response.Count)
	assert.Equal(t, 1, len(response.Items))
	assert.Equal(t, id, response.Items[0].ID)
}

func Test_GetGalleries(t *testing.T) {
	client := New(Cosmo, *key)

	time.Sleep(time.Second)
	response, err := client.Get(GalleriesAPI, Request{
		Title: "sex",
	})
	assert.Nil(t, err)
	assert.Equal(t, DefaultLimit, response.Count)
	assert.Equal(t, DefaultLimit, len(response.Items))

	// get by ID
	id := response.Items[0].GroupID
	time.Sleep(time.Second)
	response, err = client.Get(GalleriesAPI, Request{
		GroupID: id,
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, response.Count)
	assert.Equal(t, 1, len(response.Items))
	assert.Equal(t, id, response.Items[0].GroupID)
}
