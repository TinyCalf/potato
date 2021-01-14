package compnents

import (
	"potato/piface"
	"fmt"
	"net"
)

// Connector 是连接器，处理客户端连接
type Connector struct {
	BaseCompnent
	App piface.IApplication
	IPVersion string
	IP string
	Port int
	OnNewSession func(socket *net.TCPConn)
}

//NewConnector 创建一个Connector
func NewConnector(app piface.IApplication) piface.IConnector {
	return &Connector {
		App: app,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
}

// GetName 获取组件名称
func (c *Connector) GetName() string {
	return "Connector"
}

// OnAppStart 在App启动时直接启动connector
func (c *Connector) OnAppStart() {
	c.Start()
}

// Start 启动服务
func (c *Connector) Start() {
	fmt.Println("Connector Starting...")

	addr, err := net.ResolveTCPAddr(c.IPVersion, 
									fmt.Sprintf("%s:%d", c.IP, c.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error")
		return
	}

	listener, err := net.ListenTCP(c.IPVersion, addr)
	if err != nil {
		fmt.Println("listen tcp failed")
		return
	}
	go func() {
		sessionService := c.App.GetComponent("SessionService").(
			piface.ISessionService)
		
		for {
			// 获取新连接
			socket, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept listener error, ", err)
				continue
			}

			//将连接交给SessionService处理
			fmt.Println("get new session", socket)
			sessionService.Create(socket)
		}
	}()
}



