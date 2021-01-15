package common

import (
	"potato/piface"
	"fmt"
	"net"
	"io"
	"sync"
)

// Session 是会话
type Session struct {
	App piface.IApplication
	ID uint32
	Property map[string]interface{}
	Socket *net.TCPConn
	closed bool //socket是否关闭
	writeChan chan []byte //写通道 无缓冲
	exitReaderChan chan bool //read routine退出通道
	exitWriterChan chan bool //write routine退出通道

	lock sync.RWMutex //对session的操作锁
}

// NewSession 创建一个新session
func NewSession(app piface.IApplication, 
				id uint32, 
				socket *net.TCPConn) piface.ISession {
	return &Session {
		App: app,
		ID: id,
		Property: make(map[string]interface{}),
		Socket: socket,
		closed: false,
		writeChan: make(chan []byte),
		exitReaderChan: make(chan bool),
		exitWriterChan: make(chan bool),
	}
}

// GetID ..
func (s *Session) GetID() uint32 {
	return s.ID
}

// Start 启动routine开始工作，收发信息
func (s *Session) Start() {
	fmt.Printf("session %d start working\n", s.GetID())
	go s.startReader()
	go s.startWriter()
}

func (s *Session) startReader() {
	fmt.Printf("Session %d reader routine running\n", s.ID)
	defer fmt.Printf("Session %d reader routine exited\n", s.ID)
	defer s.App.GetComponent("SessionService").
		(piface.ISessionService).Del(s.ID)

	for {
		select {
		case <-s.exitReaderChan:
			return
		default:
			//读取新消息，阻塞直到读取完整的消息
			msg, err:= s.readMessage()
			if err != nil {
				return
			}

			fmt.Printf("Session %d get new msg: %v\n", s.ID, msg)

			//将消息传递给handlerservice处理
			hs := s.App.GetComponent("HandlerService").(piface.IHandlerService)
			go hs.DoHandle(s, msg)
		}	
	}
}

func (s *Session) readMessage() (piface.IMessage, error) {
	unpacker := NewMessageUnpacker()
	msg := NewMessage()

	//读取msg head
	headData := make([]byte, unpacker.GetHeadLen())
	if _, err := io.ReadFull(s.Socket, headData); err != nil {
		fmt.Println("read message head error: ", err)
		return nil, err
	}
	if err := unpacker.UnpackHead(msg, headData); err != nil {
		fmt.Println("unpack head error ", err)
	}

	//读取msg body
	bodyData := make([]byte, msg.GetLen())
	if _, err := io.ReadFull(s.Socket, bodyData); err != nil {
		fmt.Println("read message body error ", err)
		return nil, err
	}
	if err := unpacker.UnpackBody(msg, bodyData); err != nil {
		fmt.Println("unpack body error ", err)
	}

	return msg, nil
}

func (s *Session) startWriter() {
	fmt.Printf("Session %d writer routine running\n", s.ID)
	defer fmt.Printf("Session %d writer routine exited\n", s.ID)
	defer s.App.GetComponent("SessionService").
		(piface.ISessionService).Del(s.ID)
	
	for {
		select {
		case <-s.exitWriterChan:
			return
		case data, ok := <-s.writeChan:
			if ok {
				//有数据要写给客户端
				if _, err := s.Socket.Write(data); err != nil {
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				fmt.Println("writeChan is Closed")
				return
			}
		}
	}
}

//TODO 属性设置加上线程安全

// Set 设置属性
func (s *Session) Set(key string, value interface{}) {
	s.Property[key] = value
}

// Get 获取属性
func (s *Session) Get(key string) interface{} {
	return s.Property[key]
}

// Del 删除属性
func (s *Session) Del(key string) {
	delete(s.Property, key)
}

// Send 发送信息
func (s *Session) Send(msg piface.IMessage) error{
	packer := NewMessagePacker()
	data, err := packer.Pack(msg)
	if err != nil {
		fmt.Println("pack message failed, ", err)
		return err;
	}
	s.writeChan<-data
	return nil
}

// Close 释放相关资源
func (s *Session) Close() {

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.closed == true {
		return
	}
	s.closed = true

	//关闭read routine
	// if _, ok := <-s.exitReaderChan; ok {
	// 	s.exitReaderChan <- true
	// }
	close(s.exitReaderChan)
	
	//关闭write routine
	// if _, ok := <-s.exitWriterChan; ok {
	// 	s.exitWriterChan <- true
	// }
	close(s.exitWriterChan)

	//释放socket
	s.Socket.Close()
}