package main

import "fmt"

func main() {
	var x int = 11
	switch {
	case x > 0:
		fmt.Println("x>0")
		fallthrough
	case x > 1:
		fmt.Println("x>1")
	case x < 0:
		fmt.Println("x<0")

	}
}
