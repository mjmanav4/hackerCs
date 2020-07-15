package client

import (
	"encoding/json"

	"github.com/golang/glog"
)

type ShopResponse struct {
	PopularTags []PopularTags `json:"popularTags"`
	ShopList    []ShopList    `json:"shopList"`
}
type ShopList struct {
	Image      string `json:"Image"`
	Name       string `json:"Name"`
	StyleCount int    `json:"styleCount"`
}

func GetShopData(shopType string) (*ShopResponse, error) {

	url := dataUrl + shopPageUrl
	url = url + "/" + shopType

	body, err := GetBody(url)
	if err != nil {

		return nil, err
	}

	var response ShopResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		glog.Errorf("Error Occured for %v while deserializing the data: Body %v Error %v", url, body, err)
		return nil, err
	}

	return &response, nil

}
