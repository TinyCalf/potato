package tcpserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"encoding/binary"
	"bytes"
)

//connection 断开以后不可复用，应当及时清除
type connection struct {
	socket *net.TCPConn
	//如果需要管理connection，可以自定id，如果不需要，设置什么都行
	id uint32

	msgChan chan []byte
	msgBuffChan chan []byte

	sync.RWMutex
	isClosed bool
	cancel context.CancelFunc

	onDataCallBack func(msg []byte)
	onCloseCallback func()
}

func newConnection(socket *net.TCPConn) *connection {
	//初始化Conn属性
	c := &connection{
		socket:    		socket,
		isClosed:    	false,
		msgChan:     	make(chan []byte),
		msgBuffChan: 	make(chan []byte, 4096),
	}

	return c
}

func (c *connection) setID(id uint32) {
	c.id = id
}

func (c *connection) onData(callback func(msg []byte)) {
	c.onDataCallBack = callback
}

func (c *connection) onClose(callback func()) {
	c.onCloseCallback = callback
}

func (c *connection) start(pctx context.Context) {
	ctx, cancel := context.WithCancel(pctx)
	c.cancel = cancel
	go c.startReader(ctx)
	go c.startWriter(ctx)
}

func (c *connection) close() {
	fmt.Println("closing connection ", c.id)

	c.Lock()
	defer c.Unlock()

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.cancel()

	c.socket.Close()

	close(c.msgChan)
	close(c.msgBuffChan)
	
	if c.onCloseCallback != nil {
		go c.onCloseCallback()
	}
}

func (c *connection) sendMsg(data []byte) error {
	c.RLock()
	if c.isClosed == true {
		c.RUnlock()
		return errors.New("connection closed when send msg")
	}
	c.RUnlock()

	fmt.Printf("sending messsage: %s\n", string(data))

	msg, err := packMessage(data)
	if err != nil {
		fmt.Println(err)
	}

	c.msgChan <- msg

	return nil
}

func (c *connection) sendBuffMsg(data []byte) error {
	c.RLock()
	if c.isClosed == true {
		c.RUnlock()
		return errors.New("Connection closed when send buff msg")
	}
	c.RUnlock()

	msg, err := packMessage(data)
	if err != nil {
		return errors.New("Pack error msg")
	}

	c.msgBuffChan <- msg

	return nil
}






func (c *connection) getRemoteAddr() net.Addr {
	return c.socket.RemoteAddr()
}

func packMessage(data []byte) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.LittleEndian, uint32(len(data))); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, data); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (c *connection) startWriter(ctx context.Context) {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.getRemoteAddr().String(), "[conn Writer exit!]")
	defer c.close()

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.socket.Write(data); err != nil {
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				return
			}
		case data, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.socket.Write(data); err != nil {
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *connection) startReader(ctx context.Context) {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println(c.getRemoteAddr().String(), "[conn Reader exit!]")
	defer c.close()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			headData := make([]byte, 4)
			if _, err := io.ReadFull(c.socket, headData); err != nil {
				fmt.Println("read msg head error ", err)
				return
			}
			var len uint32
			dataBuff := bytes.NewReader(headData)
			if err := binary.Read(dataBuff, binary.LittleEndian, &len); err != nil {
				fmt.Println(err)
				return
			}

			var data []byte
			if len > 0 {
				data = make([]byte,len)
				if _, err := io.ReadFull(c.socket, data); err != nil {
					fmt.Println("read msg data error ", err)
					return
				}
			}

			go c.onDataCallBack(data)
		}
	}
}