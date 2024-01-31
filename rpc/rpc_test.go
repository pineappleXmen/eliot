package rpc

import (
	"fmt"
	"net"
	"testing"
)

func TestClient_Call(t *testing.T) {
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

func TestServer_Start(t *testing.T) {
	host := "192.168.31.224"
	port := "12345"
	methods := []string{"AddServiceImpl.Add"}
	services := []Service{new(AddServiceImpl)}
	server := NewServer(host, port, methods, services)
	server.Start()
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

func TestIP(t *testing.T) {
	hostname, err := net.LookupHost("localhost")
	fmt.Println(hostname, err)
}
