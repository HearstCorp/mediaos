package mediaos

import (
	"encoding/json"
)

var MediaOs = &mediaOsAPI{}

type Client interface {
	GetContent(publication PubData, key string, endpoint Endpoint, req Request) (res ContentResponse, err error)
	GetImages(publication PubData, key string, req Request) (res ImageResponse, err error)
}

type mediaOsAPI struct{}

func (m *mediaOsAPI) GetContent(publication PubData, key string, endpoint Endpoint, req Request) (res ContentResponse, err error) {
	addRequestContext(publication, key, &req)

	bytes, err := doAPICall(endpoint, req)
	if err != nil {
		return res, err
	}

	switch publication.GetApiVersion() {
	case API_V1:
		err = json.Unmarshal(bytes, &res)
		if err != nil {
			return res, err
		}
	case API_V2:
		if 0 < req.ID || 0 < req.GroupID {
			cr := ContentResponse2{}

			err = json.Unmarshal(bytes, &cr)
			if err != nil {
				return res, err
			}

			res = ContentResponse{}
			res.Count = 1
			res.Items = append(res.Items, cr.Data.toContent())
		} else {
			cr := ContentResponses2{}

			err = json.Unmarshal(bytes, &cr)
			if err != nil {
				return res, err
			}

			res = cr.toContentResponse()
		}
	}

	return res, nil
}

func (m *mediaOsAPI) GetImages(publication PubData, key string, req Request) (res ImageResponse, err error) {
	addRequestContext(publication, key, &req)

	bytes, err := doAPICall(ImagesAPI, req)
	if err != nil {
		return res, err
	}

	switch publication.GetApiVersion() {
	case API_V1:
		err = json.Unmarshal(bytes, &res)
		if err != nil {
			return res, err
		}
	case API_V2:
		ir := ImageResponse2{}
		err = json.Unmarshal(bytes, &ir)
		if err != nil {
			return res, err
		}

		res = ir.toImageResponse()
	}
	
	return res, nil
}

func addRequestContext(publication PubData, key string, req *Request) {
	req.key = key
	req.publication = publication
}
