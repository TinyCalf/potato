package rpc

import (
	"fmt"

	"github.com/TinyCalf/potato/internal/tcpserver"
)

type (
	iacceptor interface {
		onData(callback onDataCallback)
		listen(port int)
		close()
	}

	tcpacceptor struct {
		svr *tcpserver.Server	
	}

	onDataCallback func([]byte) ([]byte, error)
)


func newTCPAcceptor() iacceptor {
	return &tcpacceptor{
		svr: tcpserver.NewServer(),
	}
}

func (a *tcpacceptor) onData(callback onDataCallback) {
	a.svr.OnData(func(connID uint32, data []byte) {
		resp, err := callback(data)
		if err != nil {
			fmt.Printf("rpc error:%v\n", err)
		}
		a.svr.Send(resp, connID)
	})
}

func (a *tcpacceptor) listen(port int) {
	a.svr.Listen(port)
}

func (a *tcpacceptor) close() {
	a.svr.Close()
}