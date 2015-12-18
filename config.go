package mediaos

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type API_VERSION int

type MediaosConfig struct {
	apiVersion  API_VERSION
	urlTemplate string
	protocol    string
}

const (
	API_V1 API_VERSION = iota + 1
	API_V2
)

const (
	HTTP  = "http"
	HTTPS = "https"

	url_template = "{protocol}://{domainPort}/api/v%d/{endpoint}"

	MEDIAOS_API_VERSION = "MEDIAOS_API_VERSION"
	MEDIAOS_HTTP_SECURE = "MEDIAOS_HTTP_SECURE"
)

var initialized = false
var configInitLock = &sync.Mutex{}

var Config *MediaosConfig

func (c *MediaosConfig) LoadMediaosConfig() {
	configInitLock.Lock()

	if initialized {
		configInitLock.Unlock()
		return
	}

	c.apiVersion = API_V1
	rawApiVersion := os.Getenv(MEDIAOS_API_VERSION)
	if "" != rawApiVersion {
		i, err := strconv.Atoi(rawApiVersion)
		if nil == err {
			switch i {
			case 1:
				c.apiVersion = API_V1
			case 2:
				c.apiVersion = API_V2
			default:
				panic(fmt.Sprintf("Invalid API version: %d", i))
			}
		}
	}

	c.urlTemplate = fmt.Sprintf(url_template, int(c.apiVersion))

	secure := os.Getenv(MEDIAOS_HTTP_SECURE)
	if "" == secure || "true" == secure || "True" == secure {
		c.protocol = HTTPS
	} else {
		c.protocol = HTTP
	}

	Config = c

	initialized = true
	configInitLock.Unlock()
}

func (c *MediaosConfig) GetApiVersion() API_VERSION {
	return c.apiVersion
}
