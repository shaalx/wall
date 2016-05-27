package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/shaalx/goutils"
	"time"
)

func main() {
	test("github.com")
	test("facebook.com")
}

func test(uri string) {
	str, err := httplib.Get(fmt.Sprintf("https://%s", uri)).SetTimeout(3*time.Second, 2*time.Second).String()
	goutils.CheckErr(err)
	fmt.Println(str)
}
