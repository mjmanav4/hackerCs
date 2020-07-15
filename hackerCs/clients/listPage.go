package client

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

type ListResponse struct {
	PopularTags []PopularTags `json:"popularTags"`
	StyleList   []StyleList   `json:"styleList"`
}
type PopularTags struct {
	TagName  string `json:"tagName"`
	TagType  string `json:"tagType"`
	TagValue string `json:"tagValue"`
}
type StyleList struct {
	CelebImage    string `json:"celebImage"`
	StyleImage    string `json:"styleImage"`
	ArticleType   string `json:"articleType"`
	Price         int    `json:"Price"`
	DiscountPrice int    `json:"discountPrice"`
}

func GetListData(shopType, name string) (*ListResponse, error) {

	url := dataUrl + listPageUrl
	url = url + "/" + shopType + "/" + name

	fmt.Sprintf("listPageUrl is %v", url)

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
