package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("--------")
	fmt.Println(strings.Join(os.Args[1:], " "))
}
