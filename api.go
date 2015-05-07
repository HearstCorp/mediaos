package mediaos

// Publication represents a publication endpoint (e.g. Cosmopolitan)
type Publication string

const (
	Cosmo Publication = "cosmopolitan"
	Elle  Publication = "elle"
)

type Endpoint string

const (
	Content      Endpoint = "content"
	Articles     Endpoint = "articles"
	Galleries    Endpoint = "galleries"
	Images       Endpoint = "images"
	Editors      Endpoint = "editors"
	Sections     Endpoint = "sections"
	Subsections  Endpoint = "subsections"
	AdCategories Endpoint = "ad_categories"
)

// Visibility flags what content states to return: published, draft, or both
type Visibility string

const (
	Published Visibility = "PUBLISHED"
	Draft     Visibility = "DRAFT"
)

// DefaultLimit is the default number of items to return if a limit is not specified
const DefaultLimit = 10

// MaxLimit is the maximum number of items that can be returned
const MaxLimit = 1000

// DateFormat is the format used for datetimes in URL params
const DateFormat = "2006-01-02T15:04"

type Client interface {
	GetContent(Request) (Response, error)
}
