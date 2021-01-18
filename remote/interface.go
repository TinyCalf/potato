package remote

import (
	"potato/piface"
)

// RPCFunc 是RPC方法，由用户实现
type RPCFunc func(string) string

// IService 是远程服务,管理所有远程节点，并提供远程调用服务
type IService interface {
	//注册其他节点信息
	RegistPeer(peerid string, host string, port int)
	//注册本节点对外提供的方法
	RegistMethod(methodName string, method RPCFunc)
	//开启该服务
	Start()
	//调用其他节点
	Call(peerid string, methodName string, req string) string
}

// ICompnent 是远程服务的组件化封装
// 比Service提供与应用相关更具体的功能
type ICompnent interface {
	piface.ICompnent
	IService
}

// iRPC 远程调用适配器接口，方便以后扩展各种第三方rpc
type iRPC interface {
	//开启rpc服务
	start(host string, port int)
	//调用远程方法
	call(host string, port int, methodName string, msg string) string
	//设置被调用时需要被执行的函数
	setOnCalled(rpcOnCalled) 
}

//rpcOnCalled rpc被调用时的回调
type rpcOnCalled func(methodName string, msg string) string