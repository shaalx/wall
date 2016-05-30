package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"github.com/everfore/rpcsv"
	"github.com/shaalx/goutils"
	"github.com/shaalx/wall/httplib"
	"io/ioutil"
	"testing"
	"time"
)

var (
	rpc_tcp_server = "tcphub.t0.daoapp.io:61142"
	// rpc_tcp_server = ":8800"
)

func TestHTTPS(t *testing.T) {
	WallLoop()
}

func https_(uri string) []byte {
	b, err := httplib.Get(fmt.Sprintf("%s", uri)).SetTimeout(3*time.Second, 2*time.Second).Bytes()
	if goutils.CheckErr(err) {
		return goutils.ToByte(err.Error())
	}
	// fmt.Println(goutils.ToString(b))
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

func WallLoop() {
	in := make([]byte, 1)
	clt := rpcsv.RPCClientWithCodec(rpc_tcp_server)
	defer clt.Close()
	out := rpcsv.Job{}
	for {
		err := clt.Call("RPC.Wall", &in, &out)
		if goutils.CheckErr(err) {
			time.Sleep(1e9)
			clt = rpcsv.RPCClientWithCodec(rpc_tcp_server)
			continue
		}
		fmt.Println("Wall-result:", out)
		b := https_(out.Target)
		Upload(b, out.Name)
		out.Result = b
		ret := make([]byte, 1)
		err = clt.Call("RPC.WallBack", &out, &ret)
		goutils.CheckErr(err)
		out.Result = nil
		time.Sleep(1e9)
	}

}
