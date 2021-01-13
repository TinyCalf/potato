package piface

// IApplication 是应用接口
type IApplication interface {
	//增加一个组件
	AddComponent(compnent ICompnent)
	//获取一个组件
	GetComponent(name string) ICompnent
	//启动服务，阻塞
	Start()
}