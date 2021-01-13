package piface

//IConnector 是连接器,同时是一个App组件
type IConnector interface {
	ICompnent
	Start()
}