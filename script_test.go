package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/shaalx/goutils"
	"testing"
	"time"
)

func TestHTTPS(t *testing.T) {
	https_("github.com")
	https_("facebook.com")
}

func https_(uri string) {
	str, err := httplib.Get(fmt.Sprintf("https://%s", uri)).SetTimeout(3*time.Second, 2*time.Second).String()
	goutils.CheckErr(err)
	fmt.Println(str)
}
