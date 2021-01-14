package piface

// IHandlerService ..
type IHandlerService interface {
	ICompnent
	AddRouter(id uint32, router IRouter)
	DoHandle(session ISession, msg IMessage)
}