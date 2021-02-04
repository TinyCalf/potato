package tcpserver

import (
	"fmt"
	"net"
	"context"
)

type acceptor struct {
	onNewConnCallback func(socket *net.TCPConn)
}

func newAcceptor(callback func(socket *net.TCPConn)) *acceptor {
	return &acceptor {
		onNewConnCallback: callback,
	}
}

func (a *acceptor) listen(ctx context.Context, port int) error{

	addr, err := net.ResolveTCPAddr("tcp4", 
		fmt.Sprintf("%s:%d", "0.0.0.0", port))
	if err != nil {
		return fmt.Errorf("resolve tcp addr error")
	}

	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		return fmt.Errorf("listen tcp failed")
	}

	go func() {
		defer listener.Close()
		fmt.Println("tcp server listenning on ", port)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				socket, err := listener.AcceptTCP()
				if err != nil {
					fmt.Println("accept listener error, ", err)
					continue
				}
				fmt.Printf("accepted new connection from %v\n", socket.RemoteAddr())
				go a.onNewConnCallback(socket)
			}
		}
	}()

	return nil
}