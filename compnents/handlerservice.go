package compnents

import (
	"potato/piface"
	"fmt"
)

// HandlerService ..
type HandlerService struct {
	BaseCompnent
	routers map[uint32]piface.IRouter
}

// NewHandlerService ..
func NewHandlerService() piface.IHandlerService {
	return &HandlerService {
		routers: make(map[uint32]piface.IRouter),
	}
}

// GetName 获取组件名称
func (hs *HandlerService) GetName() string {
	return "HandlerService"
}

// AddRouter ..
func (hs *HandlerService) AddRouter(id uint32, router piface.IRouter) {
	if id == 0 {
		panic("不能设置routeID为0")
	}
	hs.routers[id] = router
}

// DoHandle ..
func (hs *HandlerService) DoHandle(session piface.ISession, 
									msg piface.IMessage) {
	routeID := msg.GetRouteID()
	router := hs.routers[routeID]

	if router == nil {
		fmt.Printf("undefined routerID %d\n", routeID)
		return
	}

	router.Before(session, msg)
	router.Handle(session, msg)
	router.After(session, msg)
}

