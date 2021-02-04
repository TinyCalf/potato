package rpc

import (
	"bytes"
	"encoding/binary"
)

type message struct {
	pkgid uint32        //消息包ID，用来标注返回结果属于哪个请求
	namespace string    //命名空间
	serviceName string	//服务名称
	methodName string	//函数方法名称
	data []byte			//消息体数据
}

// Unpack 拆包
// [Len][NameSpace][Len][Sid][Len][Data]
func unpack(data []byte) (*message, error) {
	msg := &message{}
	buf := bytes.NewBuffer(data)

	number, err := readUint32(buf)
	if err != nil {
		return nil, err
	}
	msg.pkgid = number

	res, err := readProperty(buf)
	if err != nil {
		return nil, err
	}
	msg.namespace = string(res)

	res, err = readProperty(buf)
	if err != nil {
		return nil, err
	}
	msg.serviceName = string(res)

	res, err = readProperty(buf)
	if err != nil {
		return nil, err
	}
	msg.methodName = string(res)

	res, err = readProperty(buf)
	if err != nil {
		return nil, err
	}
	msg.data = res

	return msg, nil
}

func pack(msg *message) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	
	if err := writeUint32(dataBuff, msg.pkgid); err != nil {
		return nil, err
	}

	if err := writeProperty(dataBuff, []byte(msg.namespace)); err != nil {
		return nil, err
	}

	if err := writeProperty(dataBuff, []byte(msg.serviceName)); err != nil {
		return nil, err
	}

	if err := writeProperty(dataBuff, []byte(msg.methodName)); err != nil {
		return nil, err
	}

	if err := writeProperty(dataBuff, msg.data); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}





func readUint32(buf *bytes.Buffer) (uint32, error) {
	var data uint32
	if err := binary.Read(buf, binary.BigEndian, 4); err != nil {
		return 0, err
	}
	return data, nil
}

func readProperty(buf *bytes.Buffer) ([]byte,error) {
	var len uint32
	if err := binary.Read(buf, binary.BigEndian, &len); err != nil {
		return nil, err
	}
	
	data := make([]byte, len)
	if err := binary.Read(buf, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func writeUint32(buf *bytes.Buffer, i uint32) error{
	return binary.Write(buf, binary.LittleEndian, i);
}

func writeProperty(buf *bytes.Buffer, i []byte) error {
	err := binary.Write(buf, binary.LittleEndian, uint32(len(i)))
	if err != nil {
		return err
	}

	err = binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		return err
	}

	return nil
}