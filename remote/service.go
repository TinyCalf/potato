package remote

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
)

type peer struct {
	id string
	host string
	port int
}

// Service 远程服务
type Service struct {
	peerid string
	host string
	port int
	peers map[string]*peer
	methods map[string](func(string) string)
}

// NewService ..
func NewService(peerid string, host string, port int) IService {
	return &Service{
		peerid: peerid,
		host:host,
		port:port,
		peers: make(map[string]*peer),
		methods: make(map[string](func(string) string)),
	}
}

// Rpccall ..
type Rpccall struct {
	service *Service
}

// Request ..
type Request struct {
	method string
	msg string
}

// Response ..
type Response struct{
	msg string
}

// Call ..
func (r *Rpccall) Call(req Request, resp *Response) error {
	method := r.service.methods[req.method]
	if method == nil {
		fmt.Println("mothod not found")
	}
	resp = &Response{method(req.msg)}	
	return nil
}

// AddPeer 增加远程节点
func (rs *Service) AddPeer(peerid string, host string, port int) {
	if _, ok := rs.peers[peerid]; ok {
		panic(fmt.Sprintf("duplicate peerid %d",peerid))
	}
	rs.peers[peerid] = &peer{
		id: peerid,
		host: host,
		port: port,
	}
}

// RegistMethod 注册方法
func (rs *Service) RegistMethod(name string, 
								method func(string) string) {
	rs.methods[name] = method
}

// Start 开启远程服务
func (rs *Service) Start() {
	rpc.Register(Rpccall{rs})
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", rs.host, rs.port))
    if err != nil {
        panic(err)
	}
    go http.Serve(lis, nil)
}

// Call 调用远程接口
func (rs *Service) Call(peerid string, name string, msg string) string{
	peer := rs.peers[peerid]
	if peer == nil {
		return ""
	}

	method := rs.methods[name]
	if method == nil {
		return ""
	}

	conn, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", peer.host, peer.port))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req := Request{method:name, msg: msg}
	var res Response

	err = conn.Call("Rpccall.Call", req, &res)
	if err != nil {
		return ""
	}

	return res.msg
}