package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	startTime := time.Now().UnixNano()
	for i := 0; i < 3; i++ {
		go exec(i, &wg)
	}
	wg.Wait()
	endTime := time.Now().UnixNano()
	used := float64(endTime-startTime) / 1e9
	fmt.Println("total used: ", used, " s")
}

func exec(x int, wg *sync.WaitGroup) {
	startTime := time.Now().UnixNano()
	length := int64(98989898)
	var res int64
	for i := int64(0); i < length; i++ {
		res += i
	}
	endTime := time.Now().UnixNano()
	used := float64(endTime-startTime) / 1e9
	fmt.Println("Thread =", x, "This num is ", res, "use time : ", used, "s")
	wg.Done()
}
