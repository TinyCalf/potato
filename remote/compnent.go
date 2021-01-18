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
	c.host = "0.0.0.0"
	c.port = 9000
	c.peers = make(map[string]*peer)
	c.methods = make(map[string](func(string) string))
	return c
}

// GetName .. 
func (c *Compnent) GetName() string{
	return "Remote"
}

// OnAppStart ..
func (c *Compnent) OnAppStart() {
	//从Config模块获取各种信息并注册
	config := app.GetComponent("Config").(config.Compnent)
	config.SetLocalPeer(appname string, peerid string)

	c.Start()
}

// SetAddress ..
func (c *Compnent) SetAddress(host string, port int) {
	c.host = host
	c.port = port
}