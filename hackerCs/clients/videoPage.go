package client

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

type VidePageResponse struct {
	PopularTags []PopularTags `json:"popularTags"`
	Videos      []Videos      `json:"videos"`
}

type VideoTags struct {
	TagValue string `json:"tagValue"`
	TagType  string `json:"tagType"`
	TagName  string `json:"tagName"`
}
type Videos struct {
	VideoID   string      `json:"videoId"`
	VideoTags []VideoTags `json:"videoTags"`
	VideoLink string      `json:"videoLink"`
}

func GetVideoData() (*VidePageResponse, error) {

	url := dataUrl + allVideosUrl

	fmt.Sprintf("listPageUrl is %v", url)

	body, err := GetBody(url)
	if err != nil {
		return nil, err
	}

	var response VidePageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		glog.Errorf("Error Occured for %v while deserializing the data: Body %v Error %v", url, body, err)
		return nil, err
	}

	return &response, nil

}
