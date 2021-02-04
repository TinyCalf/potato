package rpc

import (
	"fmt"
	"sync"
	"time"

	"github.com/TinyCalf/potato/internal/tcpserver"
)

type (
	imailbox interface {
		connect()
		call(message) (message, error)
		close()
	}

	tcpmailbox struct {
		remoteHost string
		remotePort int
		cli *tcpserver.Client
		requests map[uint32]chan*message
		requestsLock sync.RWMutex
	}
)

func newTCPMailBox (remoteHost string, remotePort int) *tcpmailbox {
	mb := &tcpmailbox{
		remoteHost: remoteHost,
		remotePort: remotePort,
		cli: tcpserver.NewClient(),
		requests: make(map[uint32]chan*message),
	}

	cli := tcpserver.NewClient()
	//设置收到消息的回调
	cli.OnData(func(data []byte) {
		defer func() {
			//主要用于恢复下面对已关闭管道的写操作
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()

		//拆包消息
		msg, err := unpack(data)
		if err != nil {
			fmt.Println("data unpack error, ", err)
			mb.close()
			return
		}

		//找到该消息pkgid对应的消息通道
		//并通知请求的结果已经返回
		mb.requestsLock.RLock()
		req := mb.requests[msg.pkgid]
		mb.requestsLock.RUnlock()
		if req == nil {
			return
		}
		req <- msg
	})

	return mb
}

func (mb *tcpmailbox) connect() {
	mb.cli.Connect(mb.remoteHost, mb.remotePort)
}

// 发起rpc调用
func (mb *tcpmailbox) call(msg *message) (*message, error) {
	data, err := pack(msg)
	if err != nil {
		return nil, err
	}

	mb.cli.Send(data)

	ch := make(chan*message)
	defer close(ch)

	mb.requestsLock.Lock()
	mb.requests[msg.pkgid] = ch
	mb.requestsLock.Unlock()

	select {
	case resp := <-ch:
		return resp, nil
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("rpc time out 5000ms")
	}
}

func (mb *tcpmailbox) close() {
	mb.cli.Close()
}