package main

import (
	"testing"
)

func TestHTTPS(t *testing.T) {
	// WallLoop()
	uri := "www.baidu.com"
	https_(uri)
	uri = "baidu.com"
	https_(uri)
	uri = "baidu.com?q=123"
	https_(uri)
	uri = "baidu.com?q=123&from=toukii"
	https_(uri)
}
