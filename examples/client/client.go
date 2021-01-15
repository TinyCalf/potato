package main 

import (
	"net"
	"time"
	"potato/common"
	"fmt"
	"io"
)

func main(){
	conn, _ := net.Dial("tcp", "0.0.0.0:8999")


	
	for {
		data := []byte("hello world")
		msg := common.NewMessage()
		msg.SetRouteID(1)
		msg.SetData(data)
		msg.SetLen(uint32(len(data)))

		fmt.Println(msg)

		packer := common.NewMessagePacker()
		buf, _ := packer.Pack(msg)

		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		unpacker := common.NewMessageUnpacker()
		resp := common.NewMessage()
	
		//读取msg head
		headData := make([]byte, unpacker.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("read message head error ", err)
			return
		}
		if err := unpacker.UnpackHead(msg, headData); err != nil {
			fmt.Println("unpack head error ", err)
		}
	
		//读取msg body
		bodyData := make([]byte, msg.GetLen())
		if _, err := io.ReadFull(conn, bodyData); err != nil {
			fmt.Println("read message body error ", err)
			return
		}
		if err := unpacker.UnpackBody(msg, bodyData); err != nil {
			fmt.Println("unpack body error ", err)
		}

		fmt.Println("Response: ", resp)

		time.Sleep(1 * time.Second)
	}
	
}