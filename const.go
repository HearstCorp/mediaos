package mediaos

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type PubData interface {
	Name() string
	Domain() string
	Port() string
	DomainAndPort() string
}

type _pubData struct {
	name   string
	domain string
	port   string
}

func (p _pubData) Name() string {
	return p.name
}

func (p _pubData) Domain() string {
	return p.domain
}

func (p _pubData) Port() string {
	return p.port
}

func (p _pubData) DomainAndPort() string {
	return fmt.Sprintf("%s:%s", p.domain, p.port)
}

var Publications map[string]PubData

const (
	MEDIAOS = "MEDIAOS"
	PORT    = "PORT"
	DOMAIN  = "DOMAIN"
)

func init() {
	/*
		Add new publications here.

		Each new publication MUST be accompanied by corresponding environment
		variables to provide domain and port information.

		Variables:
		 - MEDIAOS_<publication>_DOMAIN
		 - MEDIAOS_<publication>_PORT

		 Example:
		 - MEDIAOS_MYMAGAZINE_DOMAIN
		 - MEDIAOS_MYMAGAZINE_PORT

		 Note: the publication names defined here must be in lower case and the
		 variables must be ALL CAPS.
	*/

	publicationsList := []string{
		"cosmopolitan",
		"seventeen",
		"elle",
		"esquire",
		"goodhousekeeping",
		"mediaosapi",
	}

	Publications = make(map[string]PubData)
	for _, publication := range publicationsList {
		upper := strings.ToUpper(publication)
		domainVarName := fmt.Sprintf("%s_%s_%s", MEDIAOS, upper, DOMAIN)
		domain := os.Getenv(domainVarName)
		if "" == domain {
			log.Printf("Missing environment variable: %s; omitting publication: %s", domainVarName, publication)
			continue
		}

		portVarName := fmt.Sprintf("%s_%s_%s", MEDIAOS, upper, PORT)
		port := os.Getenv(portVarName)
		if "" == port {
			log.Printf("Missing environment variable: %s; omitting publication", portVarName, publication)
			continue
		}

		Publications[publication] = _pubData{publication, domain, port}
	}
}

// Endpoint represent a distinct REST endpoint in the API
type Endpoint string

const (
	// ContentAPI content endpoint
	ContentAPI Endpoint = "content"
	// ArticlesAPI articles endpoint
	ArticlesAPI Endpoint = "articles"
	// GalleriesAPI galleries endpoint
	GalleriesAPI Endpoint = "galleries"
	// ImagesAPI images endpoint
	ImagesAPI Endpoint = "images"
	// EditorsAPI editors endpoint
	EditorsAPI Endpoint = "editors"
	// SectionsAPI sections endpoint
	SectionsAPI Endpoint = "sections"
	// SubsectionsAPI subsections endpoint
	SubsectionsAPI Endpoint = "subsections"
	// AdCategoriesAPI ad_categories endpoint
	AdCategoriesAPI Endpoint = "ad_categories"
)

// Visibility flags what content states to return: published or draft
type Visibility string

const (
	// Published visibility type
	Published Visibility = "PUBLISHED"
	// Draft visibility type
	Draft Visibility = "DRAFT"
)

// DefaultLimit is the default number of items to return if a limit is not specified
const DefaultLimit = 10

// MaxLimit is the maximum number of items that can be returned
const MaxLimit = 1000

// DateFormat is the format used for datetimes in URL params
const DateFormat = "2006-01-02T15:04"
