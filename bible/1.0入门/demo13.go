package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//make 创建空map map存储了键值的集合
	//Go map == Java HashMap == Python dict == Lua table 通常使用hash实现
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			//Printf有一大堆这种转换，Go程序员称之为动词（verb）。下面的表格虽然远不是完整的规范，但展示了可用的很多特性：
			//
			//%d          十进制整数
			//%x, %o, %b  十六进制，八进制，二进制整数。
			//%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
			//%t          布尔：true或false
			//%c          字符（rune） (Unicode码点)
			//%s          字符串
			//%q          带双引号的字符串"abc"或带单引号的字符'c'
			//%v          变量的自然形式（natural format）
			//%T          变量的类型
			//%%          字面上的百分号标志（无操作数）
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
