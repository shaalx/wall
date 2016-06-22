package main

import (
	"fmt"
	"github.com/everfore/rpcsv"
	"github.com/shaalx/goutils"
	"github.com/shaalx/wall/httplib"
	"io/ioutil"
	"time"

	"net/rpc"
	"net/url"
	"strings"
)

var (
	rpc_tcp_server = "rpchub.t0.daoapp.io:61142"
	// rpc_tcp_server = "127.0.0.1:8800"
)

func main() {
	WallLoop()
}

func WallLoop() {
	in := make([]byte, 1)
	clt := rpcsv.RPCClient(rpc_tcp_server)
	defer clt.Close()
	job := rpcsv.Job{}
	ok := false
	fmt.Print("$: ")
	for {
		fmt.Print(".")
		if ok, clt = checkNilThenReLoop(clt, false); ok {
			continue
		}
		err := clt.Call("RPC.Wall", &in, &job)
		if err != nil {
			if strings.Contains(err.Error(), "nil-job") {
				time.Sleep(5e8)
			} else {
				fmt.Println(err)
				_, clt = checkNilThenReLoop(clt, true)
			}
			continue
		}
		fmt.Println("Wall-result:", job)
		b := https_(job.Target)
		Upload(b, job.Name)
		job.Result = b
		ret := make([]byte, 1)
		// clt = rpcsv.RPCClient(rpc_tcp_server)
		err = clt.Call("RPC.WallBack", &job, &ret)
		job.Result = nil
		fmt.Println("Wall-result:", job)
		if goutils.CheckErr(err) {
			_, clt = checkNilThenReLoop(clt, true)
			clt.Call("RPC.JustJob", &job, &b)
		}
		time.Sleep(5e8)
	}

}

func waiting(i int) {
	m := i % 4
	switch m {
	case 0:
		fmt.Print("\b—")
		break
	case 1:
		fmt.Print("\b\b/")
		break
	case 2:
		fmt.Print("\b|")
		break
	case 3:
		fmt.Print("\b\\")
		break
	}
}

// 是否重新开始循环
func checkNilThenReLoop(clt *rpc.Client, reconnect bool) (bool, *rpc.Client) {
	if clt == nil || reconnect {
		clt = rpcsv.RPCClient(rpc_tcp_server)
		time.Sleep(1e9)
		fmt.Print("-")
		return true, clt
	}
	return false, clt
}

func https_(uri string) []byte {
	if !strings.HasPrefix(uri, "http") {
		uri = fmt.Sprintf("https://google.com/search?q=%s", uri)
	}
	uri_, _ := url.Parse(uri)
	qry := uri_.Query()
	qry.Add("ie", "UTF-8")
	qry.Add("sourceid", "chrome")
	qry.Add("aqs", "chrome..69i57j69i60l4.1517j0j4")
	uri_.RawQuery = qry.Encode()
	uri = uri_.String()
	// fmt.Println(uri)
	// return nil
	b, err := httplib.Get(uri).SetTimeout(3*time.Second, 2*time.Second).Bytes()
	if goutils.CheckErr(err) {
		return goutils.ToByte(err.Error())
	}
	// fmt.Println(goutils.ToString(b))
	return b
}

func Upload(bs []byte, page string) (err error) {
	req := httplib.Put(fmt.Sprintf("http://upload.daoapp.io/upload/.forbidden/%s.html", page))
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
