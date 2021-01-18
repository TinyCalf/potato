package remote

import (
	"fmt"
)

type peer struct {
	id string
	host string
	port int
}

// Service 远程服务
type Service struct {
	rpc iRPC
	peerid string
	host string
	port int
	peers map[string]*peer
	methods map[string]RPCFunc
}

// NewService ..
func NewService(peerid string, host string, port int) IService {
	s := &Service{
		peerid: peerid,
		host:host,
		port:port,
		peers: make(map[string]*peer),
		methods: make(map[string]RPCFunc),
	}

	gorpc := newGoRPC()
	gorpc.setOnCalled(func(methodName string, msg string) string {
		method := s.methods[methodName]
		if method == nil {
			return ""
		}
		return method(msg)
	})
	s.rpc = gorpc

	return s
}

// RegistPeer 注册节点信息
func (rs *Service) RegistPeer(peerid string, host string, port int) {
	if _, ok := rs.peers[peerid]; ok {
		panic(fmt.Sprintf("duplicate peerid %s",peerid))
	}
	rs.peers[peerid] = &peer{
		id: peerid,
		host: host,
		port: port,
	}
}

// RegistMethod 注册本节点对外提供的方法
func (rs *Service) RegistMethod(methodName string, 
								method RPCFunc) {
	rs.methods[methodName] = method
}

// Start 开启远程服务
func (rs *Service) Start() {
	rs.rpc.start(rs.host, rs.port)
}

// Call 调用远程接口
func (rs *Service) Call(peerid string, methodName string, msg string) string{
	peer := rs.peers[peerid]
	if peer == nil {
		return ""
	}

	return rs.rpc.call(peer.host, peer.port, methodName, msg)
}