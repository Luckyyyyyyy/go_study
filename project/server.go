package main

import (
	"fmt"
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
	//当前链接的业务
	fmt.Println("连接成功")
}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listen.Close()
	for {
		//accept
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Listen Accept err:", err)
			continue
		}

		//do handler 开启一个新线程
		go this.Handler(conn)
	}

	//close listen socket
}
