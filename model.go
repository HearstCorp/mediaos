package mediaos

import "time"

type ContentType string

const (
	ArticleType    ContentType = "article"
	GalleryType    ContentType = "gallery"
	ImageType      ContentType = "image"
	SectionType    ContentType = "section"
	SubsectionType ContentType = "subsection"
	EditorType     ContentType = "editor"
	AdCategoryType ContentType = "ad_category"
)

type Request struct {
	Title         string
	Slug          string
	ContentTypeID []int
	SectionID     []int
	SubsectionID  []int
	CollectionID  []int
	AdCategoryID  []int
	Editor1       []int
	Editor2       []int
	Editor3       []int
	PublishedFrom time.Time
	PublishedTo   time.Time
	Visibility    Visibility
	IgnoreCache   bool
	AllImages     bool
	GetImageCuts  bool
	HDWActive     bool
	OrderBy       string
	Start         int
	Limit         int
	JsonpCallback string

	key         string
	publication Publication
}

type Response struct {
	Count int       `json:"count"`
	Items []Content `json:"items"`
}

type Content struct {
	ID     int         `json:"content_id"`
	Type   ContentType `json:"resource_type"`
	Title  string      `json:"title"`
	Images []Image     `json:"images"`
	URL    string      `json:"url"`
}

type Image struct {
	ID     int                 `json:"id"`
	UUID   string              `json:"uuid"`
	Width  int                 `json:"width"`
	Height int                 `json:"height"`
	URL    string              `json:"url"`
	Cuts   map[string]ImageCut `json:"image_cuts"`
}

type ImageCut struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}
