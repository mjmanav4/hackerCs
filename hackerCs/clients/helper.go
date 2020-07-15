package client

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func GetBody(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	var resp *http.Response
	var err error
	resp, err = dataClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func PostBody(url string, payload []byte) ([]byte, error) {

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")

	res, err := dataClient.Do(req)
	if err != nil {
		glog.Errorf("Exception occured while posting review to fkApi %v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
