package mediaos

// Publication represents a publication endpoint (e.g. Cosmopolitan)
type Publication string

const (
	// MediaOS publication
	MediaOs Publication = "mediaos-api"
	// Cosmo cosmopolitan publication
	Cosmo Publication = "cosmopolitan"
	// Elle publication
	Elle Publication = "elle"
	// Seventeen publication
	Seventeen Publication = "seventeen"
	// Good House Keeping publication
	GoodHouseKeeping Publication = "goodhousekeeping"
	// Esquire publication
	Esquire Publication = "esquire"
)

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
