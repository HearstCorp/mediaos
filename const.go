package mediaos

import (
	"fmt"
	"os"
	"strings"
)

type PubData interface {
	Name() string
	MosDomain() string
	MosPort() string
	MosDomainAndPort() string
	RamsDomain() string
	DisplayName() string
	NotificationAlias() string
	IUDomain() string
}

type _pubData struct {
	name              string
	mosDomain         string
	mosPort           string
	ramsDomain        string
	displayName       string
	notificationAlias string
	iuDomain          string
}

func (p _pubData) Name() string {
	return p.name
}

func (p _pubData) MosDomain() string {
	return p.mosDomain
}

func (p _pubData) MosPort() string {
	return p.mosPort
}

func (p _pubData) MosDomainAndPort() string {
	return fmt.Sprintf("%s:%s", p.mosDomain, p.mosPort)
}

func (p _pubData) RamsDomain() string {
	return p.ramsDomain
}

func (p _pubData) DisplayName() string {
	return p.displayName
}

func (p _pubData) NotificationAlias() string {
	return p.notificationAlias
}

func (p _pubData) IUDomain() string {
	return p.iuDomain
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

		_pubData.name is the publication name provided by RAMS.
		_pubData.ramsDomain is the consumer facing domain name

		For publications with multiple representations, add an appropriate entry
		to the publicationsAliases map, using the _pubData.name field as the key

		Each new publication MUST be accompanied by corresponding environment
		variables to provide domain and port information.

		Variables:
		 - MEDIAOS_<canonical form>_DOMAIN
		 - MEDIAOS_<canonical form>_PORT

		 Example:
		 - MEDIAOS_COSMO_DOMAIN
		 - MEDIAOS_COSMO_PORT

		 Note: the publication names defined here must be in lower case and the
		 variables must be ALL CAPS.
	*/

	Publications = make(map[string]PubData)

	v2PublicationsList := []_pubData{
		_pubData{name: "cosmopolitan_us", iuDomain: "cosmopolitan", displayName: "Cosmopolitan US", notificationAlias: "cosmo_us"},
	}

	for _, p := range v2PublicationsList {
		Publications[p.name] = p
	}

	publicationsList := []_pubData{
		_pubData{name: "cosmo", ramsDomain: "cosmopolitan", displayName: "Cosmopolitan", notificationAlias: "cosmo"},
		_pubData{name: "seventeen", ramsDomain: "seventeen", displayName: "Seventeen", notificationAlias: "seventeen"},
		_pubData{name: "elle", ramsDomain: "elle", displayName: "Elle", notificationAlias: "elle"},
		_pubData{name: "esquire", ramsDomain: "esquire", displayName: "Esquire", notificationAlias: "esquire"},
		_pubData{name: "goodhousekeeping", ramsDomain: "goodhousekeeping", displayName: "Good Housekeeping", notificationAlias: "ghk"},
		_pubData{name: "mediaos", ramsDomain: "mediaos", displayName: "Media OS", notificationAlias: "mediaos"},
		_pubData{name: "harpersbazaar", ramsDomain: "harpersbazaar", displayName: "HarpersBAZAAR", notificationAlias: "harpersbazaar"},
		_pubData{name: "housebeautiful", ramsDomain: "housebeautiful", displayName: "House Beautiful", notificationAlias: "housebeautiful"},
		_pubData{name: "countryliving", ramsDomain: "countryliving", displayName: "Country Living", notificationAlias: "countryliving"},
		_pubData{name: "popularmechanics", ramsDomain: "popularmechanics", displayName: "Popular Mechanics", notificationAlias: "popmech"},
		_pubData{name: "delish", ramsDomain: "delish", displayName: "Delish", notificationAlias: "delish"},
		_pubData{name: "marieclaire", ramsDomain: "marieclaire", displayName: "Marie Claire", notificationAlias: "marieclaire"},
		_pubData{name: "redbook", ramsDomain: "redbook", displayName: "Redbook", notificationAlias: "redbook"},
	}

	publicationsAliases := make(map[string]string)
	publicationsAliases["cosmopolitan"] = "cosmo"
	publicationsAliases["mediaosapi"] = "mediaos"
	publicationsAliases["mediaos-api"] = "mediaos"

	for _, p := range publicationsList {
		upper := strings.ToUpper(p.name)
		domainVarName := fmt.Sprintf("%s_%s_%s", MEDIAOS, upper, DOMAIN)
		p.mosDomain = os.Getenv(domainVarName)
		if "" == p.mosDomain {
			log.Warn("Missing environment variable: %s; omitting publication: %s", domainVarName, p.name)
			continue
		}

		portVarName := fmt.Sprintf("%s_%s_%s", MEDIAOS, upper, PORT)
		p.mosPort = os.Getenv(portVarName)
		if "" == p.mosPort {
			log.Warn("Missing environment variable: %s; omitting publication", portVarName, p.name)
			continue
		}

		Publications[p.name] = p
	}

	for alias, name := range publicationsAliases {
		p, ok := Publications[name]
		if ok {
			Publications[alias] = p
		}
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
