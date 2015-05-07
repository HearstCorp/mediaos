package mediaos

import (
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_setStringParam(t *testing.T) {
	params := url.Values{}

	// canary set
	setStringParam("foo", "bar", &params)
	assert.Equal(t, "bar", params.Get("foo"))

	// update
	setStringParam("foo", "baz", &params)
	assert.Equal(t, "baz", params.Get("foo"))

	// clear
	setStringParam("foo", "", &params)
	_, ok := params["foo"]
	assert.False(t, ok)
}

func Test_setIntListParam(t *testing.T) {
	params := url.Values{}

	// nil test
	setIntListParam("foo", nil, &params)
	_, ok := params["foo"]
	assert.False(t, ok)
	_, ok = params["foo:in"]
	assert.False(t, ok)

	// empty test
	setIntListParam("foo", []int{}, &params)
	_, ok = params["foo"]
	assert.False(t, ok)
	_, ok = params["foo:in"]
	assert.False(t, ok)

	// single value
	setIntListParam("foo", []int{42}, &params)
	assert.Equal(t, "42", params.Get("foo"))
	_, ok = params["foo:in"]
	assert.False(t, ok)

	// multiple values
	setIntListParam("foo", []int{42, 86, 108, 666}, &params)
	assert.Equal(t, "42,86,108,666", params.Get("foo:in"))
	_, ok = params["foo"]
	assert.False(t, ok)

	// and clear it
	setIntListParam("foo", nil, &params)
	_, ok = params["foo"]
	assert.False(t, ok)
	_, ok = params["foo:in"]
	assert.False(t, ok)
}

func Test_setDateParam(t *testing.T) {
	params := url.Values{}
	date := time.Now()

	// canary set
	setDateParam("foo", time.Now(), &params)
	assert.Equal(t, date.UTC().Format(DateFormat), params.Get("foo"))

	// clear, zero date
	setDateParam("foo", time.Time{}, &params)
	_, ok := params["foo"]
	assert.False(t, ok)
}

func Test_setVisibilityParam(t *testing.T) {
	params := url.Values{}

	// published
	setVisibilityParam(Published, &params)
	assert.Equal(t, "1", params.Get("visibility"))

	// draft
	setVisibilityParam(Draft, &params)
	assert.Equal(t, "0", params.Get("visibility"))

	// default
	setVisibilityParam("florb", &params)
	assert.Equal(t, "1", params.Get("visibility"))
}

func Test_setBooleanParam(t *testing.T) {
	params := url.Values{}

	// false omits param entirely
	setBooleanParam("foo", false, &params)
	_, ok := params["foo"]
	assert.False(t, ok)

	// true sets param
	setBooleanParam("foo", true, &params)
	assert.Equal(t, "1", params.Get("foo"))

	// false clears a previously true param
	setBooleanParam("foo", false, &params)
	_, ok = params["foo"]
	assert.False(t, ok)
}

func Test_setPaginationParams(t *testing.T) {
	params := url.Values{}

	// canary
	setPaginationParams(0, 100, &params)
	_, ok := params["start"]
	assert.False(t, ok)
	assert.Equal(t, "100", params.Get("limit"))

	// too low
	setPaginationParams(-1, 0, &params)
	_, ok = params["start"]
	assert.False(t, ok)
	assert.Equal(t, strconv.Itoa(DefaultLimit), params.Get("limit"))

	// set start > 0, limit too high
	setPaginationParams(13, MaxLimit+1, &params)
	assert.Equal(t, "13", params.Get("start"))
	assert.Equal(t, strconv.Itoa(DefaultLimit), params.Get("limit"))
}
