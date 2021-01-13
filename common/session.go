package common

import (
	"potato/piface"
	"fmt"
	"net"
)

// Session 是会话
type Session struct {
	ID uint32
	Property map[string]interface{}
	Socket *net.TCPConn
	Closed bool //socket是否关闭
}

// NewSession 创建一个新session
func NewSession(id uint32, socket *net.TCPConn) piface.ISession {
	return &Session {
		ID: id,
		Property: make(map[string]interface{}),
		Socket: socket,
		Closed: false,
	}
}

// GetID ..
func (s *Session) GetID() uint32 {
	return s.ID
}

func (s *Session) startReader() {
	//...TODO
}

func (s *Session) startWriter() {
	//...TODO
}

// Start 启动routine开始工作，收发信息
func (s *Session) Start() {
	fmt.Printf("session %d 开始收发工作", s.GetID())
	go s.startReader()
	go s.startWriter()
}

// Set 设置属性
func (s *Session) Set(key string, value interface{}) {
	s.Property[key] = value
}

// Get 获取属性
func (s *Session) Get(key string) interface{} {
	return s.Property[key]
}

// Del 删除属性
func (s *Session) Del(key string) {
	delete(s.Property, key)
}

// Send ..
func (s *Session) Send(msg string) {

}