package rpc

import (
	"fmt"
	"testing"
	"time"
)

func TestClient_Call(t *testing.T) {

	server := Server{addr: "localhost", port: "5678"}
	go server.Start()
	time.Sleep(time.Second)
	msg := RPCMsg{
		Addr:   "localhost:5678",
		Method: "Arith.Multiply",
		Args:   Args{A: 8, B: 4},
		Reply:  Reply{},
	}
	client := Client{addr: "localhost", port: "1234"}
	rpcMsg := client.Call(msg)
	fmt.Println(rpcMsg.Reply)
}
