// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//加入前缀判断
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//输出http code
		fmt.Printf("\nhttp status code :%s\n", resp.Status)

		//b, err := ioutil.ReadAll(resp.Body)
		//避免申请一个缓冲区, 直接到标准输出流
		io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
		//输出http code
		fmt.Printf("\nhttp status code :%s\n", resp.Status)
		//fmt.Printf("%s",body)
	}
}

//!-
