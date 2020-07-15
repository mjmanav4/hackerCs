package client

import (
	"encoding/json"
	"unicode"

	"github.com/golang/glog"
)

type TagsRequesPayload struct {
	tags []string `json:"tags"`
}

func IsTwoDim(tags string) int {
	for i, val := range tags {

		isup := unicode.IsUpper(val)
		if isup {
			return i
		}

	}
	return 0
}

func GetListOfTags(tags string) []string {
	tagList := make([]string, 0)

	index := IsTwoDim(tags)
	if index == 0 {
		tagList = append(tagList, tags)
		return tagList
	}

	word1 := tags[:index]
	word2 := tags[index:]

	tagList = append(tagList, word1)
	tagList = append(tagList, word2)

	return tagList
}

func GetShopByVideoTags(tags string) (*ListResponse, error) {

	url := dataUrl + tagsUrl

	tagList := GetListOfTags(tags)

	reqPayload := TagsRequesPayload{
		tags: tagList,
	}
	payload, _ := json.Marshal(reqPayload)

	body, err := PostBody(url, payload)
	if err != nil {
		return nil, err
	}

	var feedList ListResponse
	err = json.Unmarshal(body, &feedList)
	if err != nil {
		glog.Errorf("Error Occured for %v while deserializing the data: Body %v Error %v", url, body, err)
		return nil, err
	}

	return &feedList, nil

}
