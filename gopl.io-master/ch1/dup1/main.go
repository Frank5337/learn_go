// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//line := input.Text()
		//counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	//map的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化。
	//这种设计是有意为之的，因为能防止程序依赖特定遍历顺序，而这是无法保证的。
	//(译注：具体可以参见这里
	//http://stackoverflow.com/questions/11853396/google-go-lang-assignment-order)
	for line, n := range counts {
		if n > 1 {
			//dup1的格式字符串中还含有制表符\t和换行符\n。
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
