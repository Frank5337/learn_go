package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//             map   [k]   v
	counts := make(map[string]int)
	//os.Args[1:] go run执行  第一个参数
	//eg:  go run .\demo13_1.go C:\Users\ThinkPad\Desktop\java.txt
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	//  k     v    v++   v>1 print
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("==== %d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
