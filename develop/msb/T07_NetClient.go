package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//tuple
	conn, _ := net.Dial("tcp", "localhost:8888")

	fmt.Println(conn)

	log.Println("connected")

	//defer 在整个函数退出之前, close掉
	defer conn.Close() //finally

	//把conn读到的数据拷贝到os.stdout上面
	io.Copy(os.Stdout, conn)

}
