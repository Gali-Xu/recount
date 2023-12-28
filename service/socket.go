package service

import (
	"bufio"
	"fmt"
	"net"
)

func Socket() {
	fmt.Println("begin....")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	//循环等待客户端连接
	for {
		//等待客户端连接
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("连接失败", err)
			continue
		}
		defer listen.Close()
		fmt.Printf("客户端ip= %v", conn.LocalAddr())
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据   1.等待客户端通过conn发送信息； 2.如果客户端没有write[发送]，则协程阻塞在此
		if err != nil {
			fmt.Println("从客户端读取失败", err)
			break
		}
		recvStr := string(buf[:n]) //n代表真正读取的数据数量
		fmt.Println("收到client端发来的数据：", recvStr)

		conn.Write([]byte(recvStr)) //发送数据
	}
}
