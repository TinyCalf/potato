package remote

import (
	"potato/piface"
)

// IService 是远程服务 
type IService interface {
	AddPeer(peerid string, host string, port int)
	RegistMethod(methodName string, method func(string) string)
	Start()
	Call(peerid string, methodName string, req string) string
}

// ICompnent ..
type ICompnent interface {
	piface.ICompnent
	IService
}