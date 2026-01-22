package main

import (
	"fmt"
	"net"
)

// 定义服务器结构体
type server struct {
	Ip   string
	Port int
}

// 创建一个server接口
func Newserver(ip string, port int) *server {
	server := &server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *server) Handler(conn net.Conn) {
	//处理业务
	fmt.Println("链接建立成功")
}

// 启动服务器接口
func (this *server) Start() {
	//启动服务器的代码
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//关闭监听器
	defer listener.Close()
	//
	for {
		//accept
		conn, err := listener.Accept()
		//处理错误
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		//do handle
		go this.Handler(conn)
	}

}
