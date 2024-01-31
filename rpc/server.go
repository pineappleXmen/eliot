package rpc

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct {
	host     string
	port     string
	methods  []string
	services []Service
}

func NewServer(host, port string, methods []string, services []Service) *Server {
	return &Server{
		host:     host,
		port:     port,
		methods:  methods,
		services: services,
	}
}

func (s *Server) Start() {
	addr := s.host + ":" + s.port
	for k, v := range s.methods {
		err := rpc.RegisterName(v, s.services[k])
		if err != nil {
			fmt.Printf("register method %+v failed ", v)
		}
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("server start listening at ", addr)
	for {
		fmt.Println(listener.Addr())
		conn, err := listener.Accept()
		fmt.Println("server accept connect from conn ", conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		rpc.ServeConn(conn)
	}
}
