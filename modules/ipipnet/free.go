package ipipnet

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var H *http.Client

func init() {
	H = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func ip(ipstr string) (result string, err error) {
	resp, err := H.Get(fmt.Sprintf("http://freeapi.ipip.net/%s", ipstr))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
