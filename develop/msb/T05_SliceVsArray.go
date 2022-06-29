package main

import "fmt"

//认为是一个简单的对象
type P struct {
	name string
}

//go语言中slice map interface channel 默认使用引用传递
func m0(v [3]int) {
	v[0] = 8
}

func m1(v []int) {
	v[0] = 8
}

//这么调用不会改变原来P的值,  go是值传递,
//这么定义的函数,  往里面传值 是把这个对象整个拷贝一份, 改的是拷贝的那份
func m2(p P) {
	p.name = "ots"
}

//明确要求传一个指针, 这么传才相当于java中的引用, 才可以修改p的值
func m3(p *P) {
	p.name = "*P" // p -> name
}

func main() {
	var v1 = []int{1, 2, 3}
	var v2 = v1

	m1(v1)
	fmt.Println(v1)

	fmt.Println(v2)

	var v3 = [3]int{1, 2, 3}
	m0(v3)
	fmt.Println(v3)

	fmt.Println(P{"Frozen"})

	var pp P = P{"Fiona"}
	m2(pp)
	fmt.Println(pp)

	m3(&pp)
	fmt.Println(pp)
}
