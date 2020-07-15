package client

import (
	"encoding/json"

	"github.com/golang/glog"
)

func GetShopByVideoId(videoId string) (*ListResponse, error) {

	url := dataUrl + videIdUrl
	url = url + "/" + videoId

	body, err := GetBody(url)
	if err != nil {

		return nil, err
	}

	var response ListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		glog.Errorf("Error Occured for %v while deserializing the data: Body %v Error %v", url, body, err)
		return nil, err
	}

	return &response, nil
}
