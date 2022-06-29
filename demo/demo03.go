package main

import "fmt"

//  找出数组中和为给定值的两个元素的下标并计算和
//  例如数组[1,3,5,8,7]， 找出两个元素之和等于8的下标分别是（0，4）和（1，2）

func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := 0; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}

func main() {
	arr := [5]int{1, 3, 5, 8, 7}
	myTest(arr, 8)
}
