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
	for {
		// b := https_("https://www.google.com/search?q=golang&oq=golang&aqs=chrome..69i57j69i60l4.1517j0j4&sourceid=chrome&ie=UTF-8")
		b := https_("https://www.google.com/search?q=Blues+Ain%27t+Never+Gonna+Die+%E7%94%B5%E5%90%89%E4%BB%96%E8%B0%B1&oq=Blues+Ain%27t+Never+Gonna+Die+%E7%94%B5%E5%90%89%E4%BB%96%E8%B0%B1&aqs=chrome..69i57.10765j0j7&sourceid=chrome&ie=UTF-8")
		// b := https_("https://github.com")
		Upload(b, "guitar")
		break
		time.Sleep(1e9)
	}
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
