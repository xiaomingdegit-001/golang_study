package main

/**
 * 轻量级线程
 */

import (
	"fmt"
	"sync"
)

// 计数器, 类似Java的CountDownLatch
var wg = sync.WaitGroup{}

func test() {
	// 计数器减1
	defer wg.Done()
	fmt.Println("hello goroutine.")
}

func main() {
	var num = 10
	wg.Add(num)
	for i := 0; i < num; i++ {
		// 使用go关键字, 启动goroutine, 执行test()函数
		go test()
	}
	fmt.Println("hello main.")
	// 阻塞等待结束
	wg.Wait()
}
