package common

type Message struct {
	Len uint32
	ID uint32
	AppID uint32
	RouteID uint32
	Data []byte
}

func (m *Message) GetLen() uint32 {
	return m.Len
}

func (m *Message) SetLen(len uint32) {
	m.Len = len
}

func (m *Message) GetID() uint32 {
	return m.ID
}

func (m *Message) SetID(id uint32) {
	m.ID = id
}

func (m *Message) GetAppID() uint32 {
	return m.AppID
}

func (m *Message) SetAppID(appID uint32) {
	m.AppID = appID
}

func (m *Message) GetRouteID() uint32 {
	return m.RouteID
}

func (m *Message) SetRouteID( routeID uint32) {
	m.RouteID = routeID
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}