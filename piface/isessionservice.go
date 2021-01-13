package piface

import "net"

// ISessionService 是会话管理器
type ISessionService interface {
	ICompnent
	Create(socket *net.TCPConn) 					//创建一个新session
	Get(sid uint32) ISession 					//通过sessionid获取session
	Del(sid uint32) 							//删除session
}