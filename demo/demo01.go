package main

import "fmt"

func main() {
	// 第2个元素是100 第4个元素是200
	c := [5]int{2: 100, 4: 200, 1: 520}

	for i := 0; i < len(c); i++ {
		//fmt.Println(c[i])
	}

	for index, value := range c {
		fmt.Print(index, "=")
		fmt.Println(value)
	}

	// 通过make 创建切片
	var slice0 []int = make([]int, 10)
	slice1 := make([]int, 10)
	slice2 := make([]int, 10)
	fmt.Println(slice0, slice1, slice2)

}
