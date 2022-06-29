package main

import "fmt"

//切片
// 需要说明, slice并不是数组或数组指针, 它通过内部指针和相关属性引用数组片段, 以实现变长方案
// 1.切片: 切片是数组的一个引用, 因此切片是引用类型, 但自身是结构体, 值拷贝传递
// 2.切片的长度可以改变, 因此切片是一个可变的数组
// 3.切片遍历方式和数组一样, 可以用len()求长度, 表示可用元素数量, 读写操作不能超过该限制
// 4.cap可以求出slice最大扩张容量, 不能超出数组限制. 0 <= len(slice) <= len(array), 其中array是slice引用的数组
// 5.切片的定义: var 变量名 []类型, 比如 var str []string var arr []int
// 6.如果 slice == nil, 那么 len , cap 结果都等于 0

//创建数组的方式

func main() {
	// 声明数组, 必须给定长度
	var s0 = [4]int{1, 2, 3, 4}

	// 声明一个切片, 默认值nil
	var s1 []int

	//声明切片
	s2 := []int{}

	//第一个参数为切片类型, 第二个参数为切片大小, 第三个参数为切片容量
	var s3 []int = make([]int, 0, 0)
	fmt.Println(s0, s1, s2, s3)

	// 初始化赋值
	var s4 = []int{1, 2, 3}
	fmt.Println(s4)

	// 从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	s6 := arr[:3]
	fmt.Println(s6)
	s6 = append(s6, 6, 7, 8, 9, 10)
	fmt.Println(s6)
}
