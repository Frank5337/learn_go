package main

import "fmt"

//通过make创建切片
func main() {
	var slice0 []int = make([]int, 10)

	slice1 := make([]int, 10)
	slice2 := make([]int, 10)

	fmt.Println(slice0, slice1, slice2)

	var slice3 = make([]int, 2, 5)
	fmt.Println(slice3)

	x := []int{2, 3, 5, 7, 11}

	//y := x[len, cap]
	//y := x[1:3]  //3 5
	y := x[3:5] //   [start:end) 包头不包尾巴
	fmt.Println(y)

	z := x[:len(x)-1]
	fmt.Println(z)

}
