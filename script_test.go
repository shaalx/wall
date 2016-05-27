package main

import (
	"fmt"
	"github.com/shaalx/goutils"
	"github.com/shaalx/wall/httplib"
	"testing"
	"time"
)

func TestHTTPS(t *testing.T) {
	// https_("github.com")
	// https_("https://www.google.com/search?q=golang&oq=golang&aqs=chrome..69i57j69i60l4.1517j0j4&sourceid=chrome&ie=UTF-8")
	// https_("https://www.google.com/url?q=https://golang.org/&sa=U&ved=0ahUKEwiFy9WqwfrMAhXjnqYKHSeKA4cQFggUMAA&usg=AFQjCNFcrPeHEGHK2GcA7xFAvhgbQGjr8Q")
	https_("https://golangnews.com")
}

func https_(uri string) {
	str, err := httplib.Get(fmt.Sprintf("%s", uri)).SetTimeout(3*time.Second, 2*time.Second).String()
	goutils.CheckErr(err)
	fmt.Println(str)
}
