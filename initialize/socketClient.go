package initialize

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func SocketClientInit() {
	conn, err := net.Dial("tcp", "0.0.0.0:27017")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer conn.Close()                       //关闭连接
	inputReader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入【终端】
	for {
		//终端读取一行用户输入，并准别发送给服务器
		input, _ := inputReader.ReadString('\n') //读取用户输入

		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "exit" { //输入q则退出
			return
		}

		//将inputInfo发送给服务器
		_, err = conn.Write([]byte(inputInfo)) //发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("读取失败", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
