package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/shaalx/goutils"
	"net/http"
	"time"
)

func main() {
	// test()
	// return
	http.HandleFunc("/http/", suffer)
	http.HandleFunc("/https/", https)
	http.ListenAndServe(":80", nil)
}

func suffer(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form, req.RequestURI)
	uri := req.RequestURI
	fmt.Println(uri[6:])
	resp := httplib.Get(fmt.Sprintf("http://%s", uri[6:])).SetTimeout(3*time.Second, 2*time.Second)
	bs, err := resp.Bytes()
	goutils.CheckErr(err)
	fmt.Println(goutils.ToString(bs))
	fmt.Fprint(rw, goutils.ToString(bs))
}

func https(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form, req.RequestURI)
	uri := req.RequestURI
	fmt.Println(uri[7:])
	resp := httplib.Get(fmt.Sprintf("https://%s", uri[7:])).SetTimeout(3*time.Second, 2*time.Second)
	bs, err := resp.Bytes()
	goutils.CheckErr(err)
	fmt.Println(goutils.ToString(bs))
	fmt.Fprint(rw, goutils.ToString(bs))
}

func test() {
	str, err := httplib.Get("http://bookmark.daoapp.io").String()
	goutils.CheckErr(err)
	fmt.Println(str)
}
