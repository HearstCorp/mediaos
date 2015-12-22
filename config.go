package mediaos

import (
	"os"
	"sync"
)

type MediaosConfig struct {
	protocol string
}

const (
	HTTP  = "http"
	HTTPS = "https"

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
