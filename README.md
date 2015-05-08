# MediaOS Golang client

mediaos is a simple Go library for interacting with the MediaOS API. 

## Installation 

`go get github.com/Hearst-DD/mediaos`

## Usage

```go

var api = mediaos.New(mediaos.Cosmo, "qu9jcw5ndadqcgna8wtw8wdc")

content, err := api.GetContent(mediaos.Request{
	Title: "sex tips",
	GetImageCuts:  true,
	// additional query params...
})

if err != nil {
	panic(err)
}

for i, item := range content.Items {
	// do something...
}

```