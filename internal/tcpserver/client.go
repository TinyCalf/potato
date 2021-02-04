package tcpserver

import (
	"net"
	"fmt"
	"sync"
	"context"
)

//Client TCP客户端
type Client struct {
	onDataCallBack func([]byte) //客户端收取到信息时的回调
	host string
	port int

	conn *connection
	closed bool
	lock sync.RWMutex
}

// NewClient 获取一个新的Client对象
func NewClient() *Client {
	return &Client{
		closed: true,
	}
}

// OnData 设置收到信息时的回调函数,不设置就会忽略收到的信息
func (c *Client) OnData(callback func([]byte)) {
	c.onDataCallBack = callback
}

// Connect 开启连接
// Client是可以复用的，关闭以后可以再开启
func (c *Client) Connect(host string, port int) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if !c.closed {
		return fmt.Errorf("client already connected")
	}

	c.host = host
	c.port = port

	addr, err:=net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("resolve tcp addr error")
		return err
	}

	socket, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("tcp client start failed")
		return err
	}

	c.closed = false

	//构建一个新的connection并启动
	c.conn = newConnection(socket)
	c.conn.onClose(func() {
		c.Close()
	})
	c.conn.onData(func(data []byte) {
		if c.onDataCallBack != nil {
			c.onDataCallBack(data)
		}else {
			fmt.Println("client received new message but found no callback")
		}	
	})
	c.conn.start(context.TODO())

	return nil
}

// Close 关闭客户端
func (c *Client) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.closed {
		return
	}
	c.closed = true

	//关闭连接
	c.conn.close()
	c.conn = nil
}

// Send 发送信息
func (c *Client) Send(data []byte) error {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.closed {
		return fmt.Errorf("closed")
	}

	return c.conn.sendMsg(data)
}


