package piface

// IHandlerService ..
type IHandlerService interface {
	AddRouter(id uint32, router IRouter)
	DoHandle(session ISession, msg IMessage)
}