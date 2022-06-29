package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 默认值是 0
	var o int
	//o = 100
	var s = 123
	ss := "asd"
	str := "qwe"
	fmt.Println("姓名\t 年龄\t 籍贯\t 住址\njohn\t 12\t\t 河北\t 北京")
	fmt.Println(o)
	fmt.Println(str)
	fmt.Println(s)
	fmt.Println(ss)
	n1, name, n3 := 100, "flk", 999
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(n1, name, n3)
	var aa = strconv.Itoa(s) + "1"
	//var bb string = strconv.Atoi("998")
	fmt.Println(a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(aa)

	var z1 int32 = 12
	var z2 int64
	var z3 int8
	z2 = int64(z1) + 127
	// int8 -128 ~ 127
	//z3 = int8(z1) + 128 X
	fmt.Println("z2=", z2, "z3=", z3)

	fmt.Println("z3=%d z3对应码值=%d", z3, z3)
	var num1 = 99
	str = fmt.Sprintf("%d", num1)
	fmt.Println(str)

}
