package remote

import (
	"potato/piface"
)

// Compnent 把Service包装成组件
type Compnent struct{
	piface.BaseCompnent
	Service
}

//NewCompnent ..
func NewCompnent() ICompnent{
	c := &Compnent{}
	c.host = "0.0.0.0"
	c.port = 9000
	c.peers = make(map[uint32]*peer)
	c.methods = make(map[string](func(string) string))
	return c
}

// GetName .. 
func (c *Compnent) GetName() string{
	return "Remote"
}

// OnAppStart ..
func (c *Compnent) OnAppStart() {
	c.Start()
}

// SetAddress ..
func (c *Compnent) SetAddress(host string, port int) {
	c.host = host
	c.port = port
}