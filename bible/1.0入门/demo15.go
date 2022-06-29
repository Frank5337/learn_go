package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

//这个程序从两个package中导入了函数，net/http和io/ioutil包，
//http.Get函数是创建HTTP请求的函数，如果获取过程没有出错，
//那么会在resp这个结构体中得到访问的请求结果。
//resp的Body字段包括一个可读的服务器响应流。
//ioutil.ReadAll函数从response中读取到全部内容；将其结果保存在变量b中。
//resp.Body.Close关闭resp的Body流，防止资源泄露，Printf函数会将结果b写出到标准输出流中。
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			fmt.Printf("Plz use http:// prefix")
			os.Exit(1)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("============================" + resp.Status + "\n")

		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//}
		//fmt.Printf("%s", b)

		//if _, err := io.Copy(dst, resp.Body); err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//}
		//fmt.Printf("%s", b)
	}
}
