package compnents

import (
	"potato/piface"
	"potato/common"
	"net"
)

// SessionService 是会话服务
type SessionService struct {
	BaseCompnent
	App piface.IApplication
	currSid uint32
	sessions map[uint32]piface.ISession
}

// NewSessionService 创建一个会话服务
func NewSessionService(app piface.IApplication) piface.ISessionService {
	return &SessionService{
		App: app,
		currSid: 0,
		sessions: make(map[uint32]piface.ISession),
	}
}

// GetApp 获取所在应用
func (ss *SessionService) GetApp() piface.IApplication {
	return ss.App
}

// GetName 获取组件名称
func (ss *SessionService) GetName() string {
	return "SessionService"
}

//Create 创建一个新session, 加入SessionService的集合
func (ss *SessionService) Create(socket *net.TCPConn) {
	session := common.NewSession(ss.currSid, socket)
	ss.sessions[ss.currSid] = session
	session.Start()
	ss.currSid ++
}

//Get 通过sessionid获取session
func (ss *SessionService) Get(sid uint32) piface.ISession {
	return ss.sessions[sid]
}

//Del 删除session
func (ss *SessionService)  Del(sid uint32) {
	delete(ss.sessions, sid)
}
						