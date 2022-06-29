package main

import "fmt"

func main() {
	// 定义一个结构体
	data0 := [5][5]struct {
		x int
	}{}

	data0[0][0].x = 100
	data0[0][1].x = 200
	data0[1][0].x = 300
	data0[1][1].x = 400
	fmt.Println(data0)
}
