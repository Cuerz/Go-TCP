package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	// 客户端可以发送单行数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入（终端）
	for {
		// 从终端读取一行输入，并发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err=", err)
		}
		// 如果用户输入 exit ，就退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		// 将line发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}

}
