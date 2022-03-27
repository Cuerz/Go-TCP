package main

import (
	"fmt"
	"net" //做网络socket开发，net包含有所有需要的函数和方法
)

func process(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close()

	for {
		// 创建新的切片
		buf := make([]byte, 1024)
		// 等待客户端通过conn发送信息
		// 如果客户端没有 Write 发送，协程阻塞在这里
		//fmt.Printf("服务器在等待客户端 %s 发送数据\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			fmt.Println("客户端退出，服务器的 Read err=", err)
			return
		}
		// 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	//tcp表示使用tcp协议，监听本地8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
	}
	defer listen.Close() //延时关闭

	//循环等待客户端连接
	for {
		//等待客户端连接
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v  客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		// 这里准备起一个协程，为客户端服务
		go process(conn)
	}
	// fmt.Println("listen suc=%v", listen)
}
