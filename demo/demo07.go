package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3, 4}

	s := data[2:4] // 34

	//读写操作的实际目标是底层数组, 只需注意索引号的差别
	s[0] += 100
	s[1] += 200

	fmt.Println(s)    //[103, 204
	fmt.Println(data) // 1,2,103,204
	fmt.Println("=================================================")
	//创建slice对象并自动分配底层数组

	//通过初始化表达式构造, 可使用索引号  索引8值为100
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))

	//使用make 创建, 指定len,cap的值
	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))

	//省略cap, 相当于cap = len
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))

	//使用make动态创建slice, 避免数组必须用常量做长度的麻烦, 还可用指针直接访问底层数组退化成底层数组操作
	ss := []int{1, 2, 3}
	// 通过& 取址符, 获取底层数组元素的指针
	p := &ss[0]

	// 通过*p 取到该元素的值, 并赋值100
	fmt.Println(p)
	*p = 100
	fmt.Println(ss)
	fmt.Println(*p)

	//至于[][]T, 是指元素类型为[]T,即切片中的元素数据类型为切片
	dataa := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(dataa)

	//直接修改struct array/slice成员
	dataaa := [5]struct {
		x int
	}{}

	sss := dataaa[:]
	dataaa[1].x = 100

	sss[2].x = 100
	fmt.Println(dataaa)
	fmt.Printf("%p, %p\n", &dataaa, &dataaa[0])

}
