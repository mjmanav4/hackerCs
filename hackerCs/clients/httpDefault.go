package client

import (
	"net/http"
	"time"
)

const (
	MaxIdleConnections  int  = 100
	KeepAliveConnection bool = true
)

// TODO make this private again after bolt is moved back here
func CreateHttpClient(RequestTimeOut int) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   !KeepAliveConnection,
			MaxIdleConns:        5,
			MaxIdleConnsPerHost: 5,
		},
		Timeout: time.Duration(RequestTimeOut) * time.Millisecond}
	return client
}
