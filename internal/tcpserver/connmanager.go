package tcpserver

import (
	"sync"
	"errors"
)

type connManager struct {
	connections map[uint32]*connection
	connLock    sync.RWMutex
	currentCid  uint32  //一个自增的id，用来创建新的connection
}

func newConnManager() *connManager {
	return &connManager{
		connections: make(map[uint32]*connection),
	}
}

func (connMgr *connManager) add(conn *connection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	conn.id = connMgr.currentCid
	connMgr.connections[conn.id] = conn
	connMgr.currentCid ++
}

func (connMgr *connManager) remove(conn *connection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	delete(connMgr.connections, conn.id)
}

func (connMgr *connManager) get(connID uint32) (*connection, error) {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	}
	return nil, errors.New("connection not found")
}

func (connMgr *connManager) len() int {
	return len(connMgr.connections)
}

// 对所有连接执行一个操作
// 用这个方法可以实现关闭所有连接
func (connMgr *connManager) execToAll(f func(c *connection)) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	for _, conn := range connMgr.connections {
		f(conn)
	}
}