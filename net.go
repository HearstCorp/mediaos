package mediaos

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	HTTP  = "http"
	HTTPS = "https"

	MEDIAOS_HTTP_SECURE = "MEDIAOS_HTTP_SECURE"

	MEDIAOS_COSMO_DOMAIN     = "MEDIAOS_COSMO_DOMAIN"
	MEDIAOS_ELLE_DOMAIN      = "MEDIAOS_ELLE_DOMAIN"
	MEDIAOS_SEVENTEEN_DOMAIN = "MEDIAOS_SEVENTEEN_DOMAIN"
	MEDIAOS_GOODHOUSEKEEPING_DOMAIN = "MEDIAOS_GOODHOUSEKEEPING_DOMAIN"
	MEDIAOS_SEVENTEEN_DOMAIN = "MEDIAOS_ESQUIRE_DOMAIN"
	MEDIAOS_MEDIAOS_DOMAIN   = "MEDIAOS_MEDIAOS_DOMAIN"

	MEDIAOS_COSMO_PORT     = "MEDIAOS_COSMO_PORT"
	MEDIAOS_ELLE_PORT      = "MEDIAOS_ELLE_PORT"
	MEDIAOS_SEVENTEEN_PORT = "MEDIAOS_SEVENTEEN_PORT"
	MEDIAOS_GOODHOUSEKEEPING_PORT = "MEDIAOS_GOODHOUSEKEEPING_PORT"
	MEDIAOS_ESQUIRE_PORT = "MEDIAOS_ESQUIRE_PORT"
	MEDIAOS_MEDIAOS_PORT   = "MEDIAOS_MEDIAOS_PORT"
)

var defaultDomain = "hearst.io"
var urlTemplate = "{protocol}://{domainPort}/api/v1/{endpoint}"
var protocol = ""

var domains map[Publication]string

func init() {
	secure := os.Getenv(MEDIAOS_HTTP_SECURE)
	if "" == secure || "true" == secure {
		protocol = HTTPS
	} else {
		protocol = HTTP
	}

	domains = make(map[Publication]string)

	populateDomains(Cosmo, MEDIAOS_COSMO_DOMAIN, MEDIAOS_COSMO_PORT)
	populateDomains(Elle, MEDIAOS_ELLE_DOMAIN, MEDIAOS_ELLE_PORT)
	populateDomains(Seventeen, MEDIAOS_SEVENTEEN_DOMAIN, MEDIAOS_SEVENTEEN_PORT)
	populateDomains(MediaOs, MEDIAOS_GOODHOUSEKEEPING_DOMAIN, MEDIAOS_GOODHOUSEKEEPING_PORT)
	populateDomains(MediaOs, MEDIAOS_ESQUIRE_DOMAIN, MEDIAOS_ESQUIRE_PORT)
	populateDomains(MediaOs, MEDIAOS_MEDIAOS_DOMAIN, MEDIAOS_MEDIAOS_PORT)
}

func populateDomains(pub Publication, domainVar, portVar string) {
	d := os.Getenv(domainVar)
	p := os.Getenv(portVar)

	if "" != d && "" != p {
		domains[pub] = fmt.Sprintf("%s:%s", d, p)
	}
}

func doAPICall(endpoint Endpoint, req Request) (result []byte, err error) {
	uri := prepareAPIUri(endpoint, req)
	log.Printf("URL: %s\n", uri)

	return doGet(uri)
}

func prepareAPIUri(endpoint Endpoint, req Request) (uri string) {
	dp := domains[req.publication]
	if "" == dp {
		dp = fmt.Sprintf("%s.%s", req.publication, defaultDomain)
	}

	uri = strings.Replace(urlTemplate, "{protocol}", protocol, 1)
	uri = strings.Replace(uri, "{domainPort}", dp, 1)
	uri = strings.Replace(uri, "{endpoint}", string(endpoint), 1)

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
