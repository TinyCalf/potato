package tcpserver

import(
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	//建立TCP服务器
	s := NewServer()
	s.OnData(func(connID uint32, data []byte) {
		s.Send(data, connID)
	})
	s.Listen(8080)

	//建立TCP客户端
	c := NewClient()
	c.Connect("0.0.0.0", 8080)
	c.Send([]byte("hello world"))
	c.OnData(func(msg []byte){
		if string(msg) != "hello world" {
			t.Fatal("Fail")
		}
	})

	time.Sleep(2*time.Second)
}