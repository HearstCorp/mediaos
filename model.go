package mediaos

import "time"

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
	Count int    `json:"count"`
	Items []Item `json:"items"`
}

type Item struct {
	Title string `json:"title"`
}
