package common

import (
	"potato/piface"
)

// HandlerService ..
type HandlerService struct {
	routers map[uint32]piface.IRouter
}

// NewHandlerService ..
func NewHandlerService() piface.IHandlerService {
	return &HandlerService {
		routers: make(map[uint32]piface.IRouter),
	}
}

// AddRouter ..
func (hs *HandlerService) AddRouter(id uint32, router piface.IRouter) {
	if id == 0 {
		panic("不能设置routeID为0")
	}
	hs.routers[id] = router
}

// DoHandle ..
func (hs *HandlerService) DoHandle(session piface.ISession, msg piface.IMessage) {
	routeID := msg.GetRouteID()
	router := hs.routers[routeID]

	router.Before(session, msg)
	router.Handle(session, msg)
	router.After(session, msg)
}