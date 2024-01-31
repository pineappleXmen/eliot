package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Client struct {
	port string
	addr string
}

type RPCMsg struct {
	Addr   string
	Method string
	Args   Args
	Reply  Reply
}

type Args struct {
	A, B int
}

type Reply struct {
	Value int
}

func (c *Client) Call(msg RPCMsg) RPCMsg {
	fmt.Println("client start sending msg ", msg)
	conn, err := net.Dial("tcp", msg.Addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return RPCMsg{Reply: Reply{Value: -1}} // 返回错误信息
	}
	defer conn.Close()
	var resultMsg RPCMsg

	// 发送请求消息
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(msg)
	if err != nil {
		fmt.Println("Error encodingoding result message:", err)
		return RPCMsg{Reply: Reply{Value: -1}} // 返回错误信息
	}

	// 读取结果消息
	decoder := gob.NewDecoder(conn)
	err = decoder.Decode(&resultMsg)
	if err != nil {
		fmt.Println("Error decoding result message:", err)
		return RPCMsg{Reply: Reply{Value: -1}} // 返回错误信息
	}
	return resultMsg
}
