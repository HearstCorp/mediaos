package mediaos

import "time"

// ResType type for enumerating different resource type values
type ResType string

const (
	// ArticleType represents an article
	ArticleType ResType = "article"
	// GalleryType represents a gallery
	GalleryType ResType = "gallery"
	// ImageType represents an image
	ImageType ResType = "image"
	// SectionType represents a section
	SectionType ResType = "section"
	// SubsectionType represents a subsection
	SubsectionType ResType = "subsection"
	// EditorType represents an editor
	EditorType ResType = "editor"
	// AdCategoryType represents an ad_category
	AdCategoryType ResType = "ad_category"
)

func GetCodeForResType(r ResType) (code string, ok bool) {
	switch r {
	case ArticleType:
		return "a", true
	case GalleryType:
		return "g", true
	}
	return
}

func GetResTypeFromCode(code string) (restype ResType) {
	switch code {
	case "a":
		return ArticleType
	case "g":
		return GalleryType
	}

	return
}

// Request is a convenient container for API request params. It is a superset of
// all valid request params across all API endpoints.
type Request struct {
	ID              int // articles
	GroupID         int // galleries, for some reason
	Title           string
	Slug            string
	ContentTypeID   []int
	SectionID       []int
	SubsectionID    []int
	CollectionID    []int
	AdCategoryID    []int
	AdCategoryIDNot []int
	AdCategoryName  string
	Editor1         []int
	Editor2         []int
	Editor3         []int
	PublishedFrom   time.Time
	PublishedTo     time.Time
	Visibility      Visibility
	IgnoreCache     bool
	AllImages       bool
	GetImageCuts    bool
	HDWActive       bool
	OrderBy         string
	Start           int
	Limit           int
	JsonpCallback   string

	key         string
	publication PubData
}

// Response encapsulates an API response
type ContentResponse struct {
	Count int       `json:"count"`
	Items []Content `json:"items"`
}

type ImageResponse struct {
	Count int     `json:"count"`
	Items []Image `json:"items"`
}

// Content represents an article or gallery (or listicle, etc.)
type Content struct {
	ID         int        `json:"id"`
	ContentID  int        `json:"content_id"`
	GroupID    int        `json:"group_id"`
	Type       ResType    `json:"resource_type"`
	Title      string     `json:"index_title"`
	Images     []Image    `json:"images"`
	URL        string     `json:"url"`
	AdCategory AdCategory `json:"ad_category"`
}

// Image represents an image object
type Image struct {
	ID     int                 `json:"id"`
	UUID   string              `json:"uuid"`
	Width  int                 `json:"width"`
	Height int                 `json:"height"`
	URL    string              `json:"url"`
	Cuts   map[string]ImageCut `json:"image_cuts"`
}

// ImageCut is a single crop of an image
type ImageCut struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

type AdCategory struct {
	AdCategoryName string `json:"ad_category_name"`
}
