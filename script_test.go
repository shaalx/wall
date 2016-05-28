package main

import (
	"fmt"
	"github.com/shaalx/goutils"
	"github.com/shaalx/wall/httplib"
	"io/ioutil"
	"testing"
	"time"
)

func TestHTTPS(t *testing.T) {
	b := https_("https://www.google.com/search?q=golang&oq=golang&aqs=chrome..69i57j69i60l4.1517j0j4&sourceid=chrome&ie=UTF-8")
	// b := https_("https://github.com")
	Upload(b, "test")
	// https_("https://www.google.com/url?q=https://golang.org/&sa=U&ved=0ahUKEwiFy9WqwfrMAhXjnqYKHSeKA4cQFggUMAA&usg=AFQjCNFcrPeHEGHK2GcA7xFAvhgbQGjr8Q")
	// https_("https://golangnews.com")
}

func https_(uri string) []byte {
	b, err := httplib.Get(fmt.Sprintf("%s", uri)).SetTimeout(3*time.Second, 2*time.Second).Bytes()
	if goutils.CheckErr(err) {
		return goutils.ToByte(err.Error())
	}
	fmt.Println(goutils.ToString(b))
	return b
}

func Upload(bs []byte, page string) (err error) {
	req := httplib.Put(fmt.Sprintf("http://upload.daoapp.io/upload/%s.html", page))
	req.Body(bs)
	resp, err := req.DoRequest()
	if goutils.CheckErr(err) {
		return
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if goutils.CheckErr(err) {
		return
	}
	fmt.Println(goutils.ToString(rb))
	return nil
}
