package feedImpl

import (
	"fmt"
	client "hackerCs/clients"
	"hackerCs/feed"
)

const (
	TagCardName       = "TAGS"
	ShopListCardName  = "SHOP_LIST"
	VideoListCardName = "VIDEO_LIST"
	ListPageCardName  = "LIST_PAGE"
)

func GetTagObject(tagType, tageValue, tagName string) feed.Tag {

	o := feed.Tag{}
	o.TagKey = tageValue
	o.TagLink = fmt.Sprintf("%s/%s/%s", dataUrl, tagType, tagName)
	return o
}

func GetTagCard(data []client.PopularTags) *feed.Tags {
	// first fill tag
	tagCard := feed.Tags{}
	tagCard.Type = TagCardName
	tagCard.Props.Title = "trending"

	for _, value := range data {

		to := GetTagObject(value.TagType, value.TagValue, value.TagName)
		tagCard.Props.TagsArray = append(tagCard.Props.TagsArray, &to)
	}

	return &tagCard
}

func ListDataConvertor(data *client.ListResponse) feed.FeedResponse {

	// first fill tag
	tagCard := GetTagCard(data.PopularTags)

	listPageCard := feed.ListPage{}
	listPageCard.Type = ListPageCardName

	for _, value := range data.StyleList {
		obj := feed.ListPagePropsArrayObject{}

		obj.CelebImage = value.CelebImage
		obj.DiscountPrice = int64(value.DiscountPrice)
		obj.Mrp = int64(value.Price)
		obj.SubTitle = value.ArticleType
		listPageCard.Props.ListArray = append(listPageCard.Props.ListArray, &obj)

	}

	feedResponse := feed.FeedResponse{}
	feedResponse.TagCard = tagCard
	feedResponse.ListPage = &listPageCard
	return feedResponse
}

func ShopPageDataConvertor(data *client.ShopResponse, shopType string) feed.FeedResponse {

	// first fill tag
	tagCard := GetTagCard(data.PopularTags)

	// fill ShopList
	// ignore t, alignmemt
	shopCard := feed.ShopList{}
	shopCard.Type = ShopListCardName
	for _, value := range data.ShopList {

		obj := feed.ShopListPropsArrayObject{}
		obj.Count = int64(value.StyleCount)
		obj.Image = value.Image
		obj.SubTitle = value.Name
		shopCard.Props.ShopArray = append(shopCard.Props.ShopArray, &obj)

	}

	feedResponse := feed.FeedResponse{}
	feedResponse.TagCard = tagCard
	feedResponse.ShopList = &shopCard

	return feedResponse

}

func VideoDataConverTor(videoData *client.VidePageResponse) feed.FeedResponse {

	// first fill tag
	tagCard := GetTagCard(videoData.PopularTags)

	videoListCard := feed.VideoList{}
	videoListCard.Type = VideoListCardName

	// fil video card
	for _, value := range videoData.Videos {
		obj := feed.VideoListPropsArrayObject{}
		obj.VideoLink = value.VideoLink

		// tod mock resposne or send video id, check with vinod
		obj.ShopLook = ""

		// video tags
		for _, t := range value.VideoTags {
			to := GetTagObject(t.TagType, t.TagType, t.TagName)
			obj.Tags = append(obj.Tags, &to)
		}

		videoListCard.Props.VideoArray = append(videoListCard.Props.VideoArray, &obj)

	}

	feedResponse := feed.FeedResponse{}
	feedResponse.TagCard = tagCard
	feedResponse.VideoList = &videoListCard
	return feedResponse
}

 
