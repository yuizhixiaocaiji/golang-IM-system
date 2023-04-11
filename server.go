package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//...当前完成的业务
	fmt.Println("连接建立成功")
}

// Start 启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		//do handler
		go this.Handler(conn)
	}

	//close listen socket
}
