/*
	go/net 提供的rpc服务实现 iRPC
*/

package remote

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
)

type gorpc struct {
	onCalled rpcOnCalled
}

func newGoRPC() iRPC {
	return &gorpc{nil}
}

func (r *gorpc) setOnCalled( f rpcOnCalled) {
	r.onCalled = f;
}

// Rpccall net/rpc指定的结构
type Rpccall struct {
	rpc *gorpc
}

// Request net/rpc指定的请求结构
type Request struct {
	MethodName string
	Msg string
}

// Response net/rpc指定的响应结构
type Response struct{
	Msg string
}

// Call 注册到Rpccall的方法
func (r *Rpccall) Call(req *Request, resp *Response) error {
	resp.Msg = r.rpc.onCalled(req.MethodName, req.Msg)
	return nil
}

//onStart rpc服务开启时需要的方法
func (r *gorpc) start(host string, port int) {
	rpc.Register(&Rpccall{r})
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
    if err != nil {
        panic(fmt.Sprintf("rpc service start err: %v", err))
	}
    go http.Serve(lis, nil)
}

//call 调用远程方法
func (r *gorpc) call(host string, port int, methodName string, msg string) string {
	conn, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req := &Request{MethodName:methodName, Msg: msg}
	res := &Response{}

	err = conn.Call("Rpccall.Call", req, res)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(res)
	return res.Msg
}