package mediaos

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type API_VERSION int

const (
	API_V1 API_VERSION = iota + 1
	API_V2
)

type PubData interface {
	Name() string
	MosDomain() string
	MosPort() string
	MosDomainAndPort() string
	RamsDomain() string
	DisplayName() string
	NotificationAlias() string
	GetApiVersion() API_VERSION
}

type _pubData struct {
	name              string
	mosDomain         string
	mosPort           string
	ramsDomain        string
	displayName       string
	notificationAlias string
	apiVersion        API_VERSION
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

func (p _pubData) GetApiVersion() API_VERSION {
	return p.apiVersion
}

var Publications map[string]PubData

const (
	MEDIAOS = "MEDIAOS"
	PORT    = "PORT"
	DOMAIN  = "DOMAIN"
	VERSION = "API_VERSION"
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

	publicationsList := []_pubData{
		_pubData{name: "cosmo", ramsDomain: "cosmopolitan", displayName: "Cosmopolitan", notificationAlias: "cosmo"},
		_pubData{name: "seventeen", ramsDomain: "seventeen", displayName: "Seventeen", notificationAlias: "seventeen"},
		_pubData{name: "elle", ramsDomain: "elle", displayName: "Elle", notificationAlias: "elle"},
		_pubData{name: "esquire", ramsDomain: "esquire", displayName: "Esquire", notificationAlias: "esquire"},
		_pubData{name: "goodhousekeeping", ramsDomain: "goodhousekeeping", displayName: "Good Housekeeping", notificationAlias: "ghk"},
		_pubData{name: "mediaos", ramsDomain: "mediaos", displayName: "Media OS", notificationAlias: "mediaos"},
		_pubData{name: "harpersbazaar", ramsDomain: "harpersbazaar", displayName: "HarpersBAZAAR", notificationAlias: "harpersbazaar"},
	}

	publicationsAliases := make(map[string]string)
	publicationsAliases["cosmopolitan"] = "cosmo"
	publicationsAliases["mediaosapi"] = "mediaos"
	publicationsAliases["mediaos-api"] = "mediaos"

	Publications = make(map[string]PubData)
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

		p.apiVersion = API_V1
		apiVersionVarName := fmt.Sprintf("%s_%s_%s", MEDIAOS, upper, VERSION)
		rawApiVersion := os.Getenv(apiVersionVarName)
		if "" != rawApiVersion {
			i, err := strconv.Atoi(rawApiVersion)
			if nil == err {
				switch i {
				case 1:
					p.apiVersion = API_V1
				case 2:
					p.apiVersion = API_V2
				default:
					panic(fmt.Sprintf("Invalid API version '%d' for %s", i, apiVersionVarName))
				}
			} else {
				panic(fmt.Sprintf("Invalid API version '%s' for %s", rawApiVersion, apiVersionVarName))
			}
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

var UrlTemplate = map[API_VERSION]string{
	API_V1: "{protocol}://{domainPort}/api/v1/{endpoint}",
	API_V2: "{protocol}://{domainPort}/v2/{endpoint}",
}

// Endpoint represent a distinct REST endpoint in the API
type Endpoint struct {
	v1 string
	v2 string
}

func (e *Endpoint) String(apiVersion API_VERSION) string {
	switch apiVersion {
	case API_V1:
		return e.v1
	case API_V2:
		return e.v2
	default:
		panic(fmt.Sprintf("Unknown ApiVersion: %d", int(apiVersion)))
	}
}

var ContentAPI = Endpoint{"content", "content"}
var ArticlesAPI = Endpoint{"articles", "content"}
var GalleriesAPI = Endpoint{"galleries", "content"}
var ImagesAPI = Endpoint{"images", "images"}
var EditorsAPI = Endpoint{"editors", ""}
var SectionsAPI = Endpoint{"sections", ""}
var SubsectionsAPI = Endpoint{"subsections", ""}
var AdCategoriesAPI = Endpoint{"ad_categories", ""}

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
