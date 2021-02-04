package tcpserver

import(
	"testing"
	"time"
)

//客户端主动断开连接，服务器应该停止其所有读写routine，从管理中删除
func TestClientDisconnect(t *testing.T) {
	//抽象出来的TCP服务很简单，只有下面几个调用

	//建立TCP服务器
	s := NewServer()
	s.OnData(func(connID uint32, data []byte) {
		//这里可以写各种连接管理、请求路由等等，这里简单写个echo
		//也可以不是一应一答的模式，比如群发，直接用s.Send(data, 1,2,3,4)就行
		s.Send(data, connID)
	})
	s.Listen(8080)

	//建立TCP客户端
	c := NewClient()
	c.Connect("0.0.0.0", 8080)

	//建立TCP客户端
	c2 := NewClient()
	c2.Connect("0.0.0.0", 8080)
	c2.Close()

	time.Sleep(2*time.Second)

	if s.connmgr.connections[1] != nil {
		t.Fatal("FAIL")
	}
}