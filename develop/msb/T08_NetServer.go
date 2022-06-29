package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")

	//错误处理
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("a client connected")

		/* go */
		go handle(conn)
	}

}

//只有一个handle函数 在死循环,  所以相当于只有一个服务
//相当于 一个餐厅一个服务员, 服务员一直在忙
//java每一个客户端一个线程, 来一个人一个线程,
//so 诞生了线程池 , 接收的一个线程池 处理的是一个线程池 用线程池来处理

//go 中 加go就完事了, 核心 其实是起了一个 go routine 起了一个协程 类比java起一个线程
// 相当于java线程池里面的任务

//给客户端写数据
func handle(conn net.Conn) {
	defer conn.Close()
	//var o int = 0
	for {
		//o+=1
		//str := "hello" + string(o)
		_, err := io.WriteString(conn, "hello \n")
		if err != nil {

		}

		//睡1s. 继续循环
		time.Sleep(1 * time.Second)
	}
}
