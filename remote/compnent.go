package remote

import (
	"potato/piface"
	"potato/config"
)

// Compnent 把Service包装成组件
type Compnent struct{
	piface.BaseCompnent
	Service
	app piface.IApplication
}

//NewCompnent ..
func NewCompnent(app piface.IApplication) ICompnent{
	c := &Compnent{}
	c.app = app
	c.peers = make(map[string]*peer)
	c.methods = make(map[string]RPCFunc)
	return c
}

// GetName .. 
func (c *Compnent) GetName() string{
	return "Remote"
}

// OnAppStart ..
func (c *Compnent) OnAppStart() {
	//从Config模块获取各种信息并注册
	config := c.app.GetComponent("Config").(config.ICompnent)

	//读取本地节点信息
	p := config.GetLocalPeerInfo()
	c.peerid = p.PeerID 
	c.host = p.Host
	c.port = p.RemotePort 

	//读取所有节点信息
	for _, pi := range config.GetAllPeerInfo() {
		c.RegistPeer(pi.PeerID, pi.Host, pi.RemotePort)
	}

	c.Start()
}

// SetAddress ..
func (c *Compnent) SetAddress(host string, port int) {
	c.host = host
	c.port = port
}