package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, World\n")
	fmt.Print("Hello, Go")
	fmt.Println()
	fmt.Print("Congratulations\n")
	fmt.Println(Add1(10, 20))
	fmt.Println(Add2(10, 15))
	fmt.Println("google" + " java" + " golang")
	fmt.Println(Add(10, 15))
	for i := 0; i < 10; i++ {
		fmt.Println("hello go;", i)
	}
}

func Add1(i int, i2 int) interface{} {
	return i + i2
}

func Add2(i int, i2 int) int {
	return i + i2
}
