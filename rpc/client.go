package rpc

import (
	"fmt"
	"log"
	"net/rpc"
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

func (c *Client) Call(msg RPCMsg) {

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}
