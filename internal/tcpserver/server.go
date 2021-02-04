package tcpserver

import(
	"net"
	"fmt"
	"context"
)

// Server TCP服务
type Server struct {
	acceptor *acceptor
	connmgr *connManager
	onDataCallback func(connID uint32, data []byte)

	ctx context.Context
	cancel context.CancelFunc
}

// NewServer 返回一个新的TCPServer
func NewServer() *Server{
	s := &Server{}

	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.connmgr = newConnManager()

	onNewConnCallback := func(socket *net.TCPConn) {
		conn := newConnection(socket)
		conn.onData(func(data []byte){
			if s.onDataCallback != nil {
				s.onDataCallback(conn.id, data)
			}
		})
		conn.onClose(func(){
			s.connmgr.remove(conn)
		})
		conn.start(s.ctx)
		s.connmgr.add(conn)
	}
	s.acceptor = newAcceptor(onNewConnCallback)

	return s
}

// OnData 设置收取到信息后的回调
func(s *Server) OnData(callback func(connID uint32, data []byte)) {
	s.onDataCallback = callback
}

// Listen TCPServer开始工作并监听端口
func(s *Server) Listen(port int) error {
	return s.acceptor.listen(s.ctx, port)
}

// Send 向一个或多个指定连接发送消息
func(s *Server) Send(data []byte, connIDs ...uint32) {
	for _, id := range connIDs {
		conn, err := s.connmgr.get(id)
		if err != nil {
			continue
		}
		if err := conn.sendMsg(data); err != nil {
			fmt.Println("sendMsg error,", err)
		}	
	}
}

// Close 安全关闭server
func(s *Server) Close() {
	//回收所有goroutine和fd
	s.cancel()
	s.connmgr.execToAll(func(conn *connection){
		s.connmgr.remove(conn)
	})
}