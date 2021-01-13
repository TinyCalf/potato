package common

import (
	"potato/piface"
)

//BaseRouter ..
type BaseRouter struct{}

//Before ..
func (br *BaseRouter) Before(session piface.ISession, msg piface.IMessage) {}

//Handle ..
func (br *BaseRouter) Handle(session piface.ISession, msg piface.IMessage) {}

//After ..
func (br *BaseRouter) After(session piface.ISession, msg piface.IMessage) {}