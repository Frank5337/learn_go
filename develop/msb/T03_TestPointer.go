package main

import "fmt"

/**
指针
java面向对象编程, go面向函数编程
go的指针是只读
*/

//v 是一个指向int类型的指针
func f(v *int) {
	*v = 8 //解引用
	//v = xx 改的是指针
	//*v =xx 改的是指针指向的部分
}

//拷贝一份v 来修改v  原来的v不变
func f2(v int) {
	v = 8
}

func main() {
	var i int = 8
	//& 取地址符
	fmt.Println(&i)

	//声明一个变量, 指向int类型的指针类型  类似java中的  Test t = new T 中的t java中的指针
	var pi *int
	//指针就是一个值, 值里面装的地址
	pi = &i
	fmt.Println(pi)

	//很多的时候用来传参数, go中传参数 只有传值1种,

	//声明 + 赋值
	m := 0
	f(&m)
	fmt.Println("before f2", m)
	f2(m)
	fmt.Println("f2: ", m)

	n := 0
	f2(n)
	fmt.Println(n)
}
