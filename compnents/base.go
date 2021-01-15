package compnents

//BaseCompnent 是基础组件，其他组件可以继承
type BaseCompnent struct {}

// GetName 获取组件名称
func (c *BaseCompnent) GetName() string {
	return "undefined"
}

// OnAppStart app启动时的回调
func (c *BaseCompnent) OnAppStart() {}

// OnAppStop app停止时的回调
func (c *BaseCompnent) OnAppStop() {}