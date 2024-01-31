package rpc

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClient_Call(t *testing.T) {
	server := Server{
		port: "1234",
		addr: "localhost",
	}
	go server.Start()
	time.Sleep(time.Second)

	dial, err := net.Dial("tcp", "192.168.31.224:1234")
	fmt.Println(dial, err)

	client := Client{
		port: "",
		addr: "",
	}
	client.Call(RPCMsg{
		Addr:   "223.74.148.163:1234",
		Method: "",
		Args:   Args{},
		Reply:  Reply{},
	})

}

func TestOne(t *testing.T) {
	// 监听端口
	listener, err := net.Listen("tcp", "192.168.31.224:12345")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on 192.168.31.224:12345...")

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 处理连接
		buffer := make([]byte, 1024)
		fmt.Println("connected!", conn)
		conn.Read(buffer)

		fmt.Println(string(buffer))
	}
}

func TestServer_Start(t *testing.T) {
	server := Server{
		port: "12345",
		addr: "192.168.31.224",
	}
	server.Start()
}

func TestIP(t *testing.T) {
	hostname, err := net.LookupHost("localhost")
	fmt.Println(hostname, err)
}
