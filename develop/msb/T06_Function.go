package main

import "fmt"

//函数可以当做变量来用,
//指向函数的指针,

//先有的函数式编程, 才有的java lambda表达式, java流式编程
func andBodySay(f func(word string)) {
	f("hello")
}

func main() {
	//f1   类型是一个函数 函数里一个参数是string
	var f1 = func(word string) {
		fmt.Println(word, "f1")
	}

	andBodySay(f1)

	var f2 = func(word string) {
		fmt.Println(word, "f2")
	}

	andBodySay(f2)

	f3 := f2
	andBodySay(f3)

	var ff func(word string)

	ff = f1
	fmt.Println(ff)
	ff("ha ha ha ha") //对应java Functional Interface

}
