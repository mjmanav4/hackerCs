package feedImpl

import (
	"context"
	client "hackerCs/clients"
	"hackerCs/feed"

	"github.com/golang/glog"
)

const (
	dataUrl = "http://localhost:8080"
)

// NewUserServer will return a refrence of userServer
func NewUserServer() *HackServer {
	u := HackServer{}
	return &u
}

// HackServer si
type HackServer struct{}

// GetUserNotes shs
func (u *HackServer) GetLandingPageData(ctx context.Context, in *feed.Request) (*feed.FeedResponse, error) {

	glog.Infof("okay")
	return &feed.FeedResponse{}, nil
}

func (u *HackServer) GetShopByVideo(ctx context.Context, in *feed.GetShopByVideoRequest) (*feed.FeedResponse, error) {

	glog.Infof("okay")
	return &feed.FeedResponse{}, nil
}

// GetUserNotes shs
func (u *HackServer) GetListPageData(ctx context.Context, in *feed.GetListPageDataRequest) (*feed.FeedResponse, error) {

	if in.GetName() == "" || in.GetType() == "" {
		return &feed.FeedResponse{}, nil
	}
	listData, err := client.GetListData(in.GetType(), in.GetName())
	if err != nil {
		glog.Errorf("db call err %v", err)
		return &feed.FeedResponse{}, nil
	}

	feedRes := ListDataConvertor(listData)

	return &feedRes, nil
}

// GetUserNotes shs
func (u *HackServer) GetAllVideos(ctx context.Context, in *feed.Request) (*feed.FeedResponse, error) {

	videoData, err := client.GetVideoData()
	if err != nil {
		glog.Errorf("db call err %v", err)
		return &feed.FeedResponse{}, nil
	}

	feedRes := VideoDataConverTor(videoData)
	if err != nil {
		glog.Errorf("convertor err %v", err)
		return &feed.FeedResponse{}, nil
	}

	glog.Infof("okay")
	return &feedRes, nil
}

// GetUserNotes shs
func (u *HackServer) GetShopPageData(ctx context.Context, in *feed.GetShopDataRequest) (*feed.FeedResponse, error) {

	// type is movie
	if in.GetType() == "" {
		glog.Errorf("err while fetching shop page")
		return &feed.FeedResponse{}, nil
	}
	shopData, err := client.GetShopData(in.GetType())

	if err != nil {
		glog.Errorf("err while fetching shop page")
		return &feed.FeedResponse{}, nil
	}

	feedRes := ShopPageDataConvertor(shopData, in.GetType())
	return &feedRes, nil
}
