package piface

// ISession 会话接口
type ISession interface{
	GetID() uint32						//获取sid
	Start() 							//让session开始接收和发送任务
	Close()								//准备回收资源
	Set(key string, value interface{}) 	//设置session属性
	Get(key string) interface{}					//获取属性
	Del(key string) 					//删除属性
	Send(msg IMessage) error				//发送消息
}