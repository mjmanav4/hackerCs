package client

import "net/http"

var dataClient *http.Client
var dataUrl string

const (
	popularVideosUr = "/popularVideos"
	shopPageUrl     = "/shop"
	listPageUrl     = "/listPage"
	allVideosUrl    = "/popularVideos"
	videIdUrl       = "/videoId"
	tagsUrl         = "/tags"
)

func InitializeIdeaClient() {
	dataClient = CreateHttpClient(100)
	dataUrl = "http://localhost:8080"
}
