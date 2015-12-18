package mediaos

import (
	"bytes"
	"net/url"
	"strconv"
	"time"
)

const (
	id = "id"
	groupId = "group_id"
)

func encodeParams(uri string, params url.Values) string {
	if API_V2 == Config.GetApiVersion() {
		key := id
		idValue := params.Get(key)
		params.Del(key)

		if "" == idValue {
			key = groupId
			idValue = params.Get(key)
			params.Del(key)
		}

		if "" != idValue {
			uri += "/" + idValue
		}
	}

	uri += "?" + params.Encode()

	return uri
}

func prepareParams(apiKey string, req Request) (params url.Values) {
	params = url.Values{}

	params.Set("_key", apiKey)

	setIntParam("id", req.ID, &params)
	setIntParam("group_id", req.GroupID, &params)
	setStringParam("title", req.Title, &params)
	setStringParam("slug", req.Slug, &params)
	setIntListParam("content_type_id", req.ContentTypeID, &params, true)
	setIntListParam("section_id", req.SectionID, &params, true)
	setIntListParam("subsection_id", req.SubsectionID, &params, true)
	setIntListParam("collection_id", req.CollectionID, &params, true)
	setIntListParam("ad_category_id", req.AdCategoryID, &params, true)
	setIntListParam("ad_category_id", req.AdCategoryIDNot, &params, false)
	setIntListParam("editor", req.Editor1, &params, true)
	setIntListParam("editor2", req.Editor2, &params, true)
	setIntListParam("editor3", req.Editor3, &params, true)
	setDateParam("publishedFrom", req.PublishedFrom, &params)
	setDateParam("publishedTo", req.PublishedTo, &params)
	setVisibilityParam(req.Visibility, &params)
	setBooleanParam("ignore_cache", req.IgnoreCache, &params)
	setBooleanParam("all_images", req.AllImages, &params)
	setBooleanParam("get_image_cuts", req.GetImageCuts, &params)
	setBooleanParam("hdw_active", req.HDWActive, &params)
	setStringParam("order_by", req.OrderBy, &params)
	setPaginationParams(req.Start, req.Limit, &params)

	return
}

func setStringParam(key, value string, params *url.Values) {
	if value == "" {
		params.Del(key)
	} else {
		params.Set(key, value)
	}
}

func setIntParam(key string, value int, params *url.Values) {
	if value == 0 {
		params.Del(key)
	} else {
		params.Set(key, strconv.Itoa(value))
	}
}

func setIntListParam(key string, value []int, params *url.Values, in bool) {
	operator := ":in"
	if !in {
		operator = ":not"
	}

	if len(value) < 1 {
		params.Del(key)
		params.Del(key + operator)
	} else {
		params.Del(key + operator)
		params.Set(key+operator, commaSeparateInts(value))
	}
}

func commaSeparateInts(input []int) string {
	var out bytes.Buffer
	for i, val := range input {
		if i > 0 {
			out.WriteRune(',')
		}
		out.WriteString(strconv.Itoa(val))
	}

	return out.String()
}

func setDateParam(key string, date time.Time, params *url.Values) {
	if date.IsZero() {
		params.Del(key)
		return
	}

	params.Set(key, date.UTC().Format(DateFormat))
}

func setVisibilityParam(viz Visibility, params *url.Values) {
	if viz == Draft {
		params.Set("visibility", "0")
	} else {
		params.Set("visibility", "1")
	}
}

func setBooleanParam(key string, value bool, params *url.Values) {
	// all pure boolean params default to false
	if value {
		params.Set(key, "1")
	} else {
		params.Del(key)
	}
}

func setPaginationParams(start, limit int, params *url.Values) {
	if start > 0 {
		params.Set("start", strconv.Itoa(start))
	} else {
		params.Del("start")
	}

	switch {
	case limit <= 0, limit > MaxLimit:
		params.Set("limit", strconv.Itoa(DefaultLimit))
	default:
		params.Set("limit", strconv.Itoa(limit))
	}
}
