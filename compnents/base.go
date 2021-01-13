package compnents

import (
	"potato/piface"
)

//BaseCompnent 是基础组件，其他组件可以继承
type BaseCompnent struct {}

//GetApp 获取所在应用
func (c *BaseCompnent) GetApp() piface.IApplication {
	return nil
}

// GetName 获取组件名称
func (c *BaseCompnent) GetName() string {
	return "undefined"
}

// OnAppStart app启动时的回调
func (c *BaseCompnent) OnAppStart() {}

// OnAppStop app停止时的回调
func (c *BaseCompnent) OnAppStop() {}