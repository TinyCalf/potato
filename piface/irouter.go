package piface

// IRouter 是路由接口
type IRouter interface {
	Before(ISession, IMessage)
	Handle(ISession, IMessage)
	After(ISession, IMessage)
}