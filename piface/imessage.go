package piface

// IMessage 是消息接口
type IMessage interface {
	GetLen() uint32
	GetID() uint32
	GetAppID() uint32
	GetRouteID() uint32
	GetData() []byte

	SetLen(uint32)
	SetID(uint32)
	SetAppID(uint32)
	SetRouteID(uint32)
	SetData([]byte)
}

// IMessagePacker 是消息打包工具
type IMessagePacker interface {
	Pack(msg IMessage) ([]byte, error)
}

// IMessageUnPacker 是消息拆包工具
type IMessageUnPacker interface {
	//获取已经指定的头部长度
	GetHeadLen()
	//拆解Head部分，head部分长度从GetHeadLen获取
	UnPackHead(msg IMessage, buf []byte) error
	//拆解Body部分
	UnPackBody(msg IMessage, buf []byte) error 
}

