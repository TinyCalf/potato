package common

import (
	"potato/piface"
	"bytes"
	"encoding/binary"
)

// MessagePacker 消息打包工具
type MessagePacker struct {}

// NewMessagePacker 获取一个消息打包工具
func NewMessagePacker() piface.IMessagePacker  {
	return &MessagePacker{}
}

// Pack 打包消息
func (mp *MessagePacker) Pack(msg piface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	//写数据长度
	err := binary.Write(buf, binary.BigEndian, msg.GetLen())
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, msg.GetID())
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, msg.GetAppID())
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, msg.GetRouteID())
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, msg.GetData())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// MessageUnpacker 拆包工具
type MessageUnpacker struct {}

// NewMessageUnpacker ..
func NewMessageUnpacker() piface.IMessageUnpacker {
	return &MessageUnpacker{}
}

// GetHeadLen 获取消息头部长度
func (m *MessageUnpacker) GetHeadLen() uint32 {
	return 4 + 4 + 4 + 4
}

// UnpackHead 拆包head
func (m *MessageUnpacker) UnpackHead(msg piface.IMessage, data []byte) error {
	buf := bytes.NewBuffer(data)

	var msgLen uint32
	if err := binary.Read(buf, binary.BigEndian, &msgLen); err != nil {
		return err
	}
	msg.SetLen(msgLen)

	var ID uint32
	if err := binary.Read(buf, binary.BigEndian, &ID); err != nil {
		return err
	} 
	msg.SetID(ID)

	var appID uint32
	if err := binary.Read(buf, binary.BigEndian, &appID); err != nil {
		return err
	} 
	msg.SetAppID(appID)

	return nil
}

// UnpackBody 拆包body
func (m *MessageUnpacker) UnpackBody(msg piface.IMessage, data []byte) error {
	buf := bytes.NewBuffer(data)
	body := make([]byte, msg.GetLen())
	if err := binary.Read(buf, binary.BigEndian, &body); err != nil {
		return err
	} 
	msg.SetData(body)

	return nil
}

// Message 是消息
type Message struct {
	Len uint32
	ID uint32
	AppID uint32
	RouteID uint32
	Data []byte
}

//NewMessage ...
func NewMessage() piface.IMessage{
	return &Message{}
}

// GetLen ..
func (m *Message) GetLen() uint32 {
	return m.Len
}

// SetLen ..
func (m *Message) SetLen(len uint32) {
	m.Len = len
}

// GetID ..
func (m *Message) GetID() uint32 {
	return m.ID
}

// SetID ..
func (m *Message) SetID(id uint32) {
	m.ID = id
}

// GetAppID ..
func (m *Message) GetAppID() uint32 {
	return m.AppID
}

// SetAppID ..
func (m *Message) SetAppID(appID uint32) {
	m.AppID = appID
}

// GetRouteID ..
func (m *Message) GetRouteID() uint32 {
	return m.RouteID
}

// SetRouteID ..
func (m *Message) SetRouteID( routeID uint32) {
	m.RouteID = routeID
}

// GetData ..
func (m *Message) GetData() []byte {
	return m.Data
}

// SetData ..
func (m *Message) SetData(data []byte) {
	m.Data = data
}