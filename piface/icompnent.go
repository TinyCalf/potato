package piface

//ICompnent 是组件的接口
type ICompnent interface {
	//获取组件名称
	GetName() string
	//app启动时的回调
	OnAppStart()
	//app停止时的回调
	OnAppStop()
}