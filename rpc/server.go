package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"net/rpc"
)

type Server struct {
	port string
	addr string
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (c *Server) Start() {
	addr := c.addr + ":" + c.port
	rpc.Register(new(HelloService))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("server start listening ", addr)
	for {
		fmt.Println(listener.Addr())
		conn, err := listener.Accept()
		fmt.Println("accpet")
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		rpc.ServeConn(conn)
	}
}

func handleRPCRequest(conn net.Conn) {
	var msg RPCMsg

	// 读取消息
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&msg)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	// 处理 RPC 请求
	resultMsg := handleRequest(msg)
	fmt.Println("server got result and sending result ", resultMsg)
	// 发送结果消息
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(resultMsg)
	if err != nil {
		fmt.Println("Error encoding result message:", err)
	}

	conn.Close()
}

func handleRequest(msg RPCMsg) RPCMsg {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return RPCMsg{Reply: Reply{Value: -1}} // 返回错误信息
	}
	defer client.Close()

	// 构造 RPC 方法调用
	err = client.Call(msg.Method, msg.Args, &msg.Reply)
	if err != nil {
		fmt.Println("Error calling", msg.Method, ":", err)
		return RPCMsg{Reply: Reply{Value: -1}} // 返回错误信息
	}
	return msg
}
