package remote

import (
	"potato/piface"
)

// IService 是远程服务 
type IService interface {
	AddPeer(peerid uint32, host string, port int)
	RegistMethod(methodName string, method func(string) string)
	Start()
	Call(peerid uint32, methodName string, req string) string
}

// ICompnent ..
type ICompnent interface {
	piface.ICompnent
	IService
	SetAddress(host string, port int)
}