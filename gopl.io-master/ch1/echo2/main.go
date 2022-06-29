// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", 1
	for index, arg := range os.Args[0:] {
		s += string(rune(sep)) + arg
		fmt.Println(sep)
		fmt.Println(index)
	}
	fmt.Println(s)
}

//!-
