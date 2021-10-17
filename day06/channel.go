package main

/**
 * 通道, 实现goroutines之间的通信
 */

import (
	"fmt"
	"math/rand"
	"time"
)

var itemChan = make(chan *item, 100)
var resultChan = make(chan *result, 100)

type item struct {
	id  int64
	num int64
}

type result struct {
	item
	sum int64
}

func producer() {
	var id int64
	for {
		id++
		num := rand.Int63n(10000)
		tmp := &item{
			id:  id,
			num: num,
		}
		itemChan <- tmp
	}
}

func consumer() {
	for tmp := range itemChan {
		sum := compute(tmp.num)
		r := &result{
			item: *tmp,
			sum:  sum,
		}
		resultChan <- r
	}
}

func startWorker(n int) {
	for i := 0; i < n; i++ {
		go consumer()
	}
}

func compute(num int64) int64 {
	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}

func printResult() {
	for r := range resultChan {
		fmt.Printf("id: %v, num: %v, sum: %v\n", r.id, r.num, r.sum)
		time.Sleep(time.Second)
	}
}

func main() {
	// 定义一个channel类型, 并生命channel内部传递数据的类型
	var ch1 chan int
	var ch2 chan string
	// channel是引用类型
	fmt.Printf("ch1: %v\n", ch1)
	fmt.Printf("ch2: %v\n", ch2)
	// 使用make函数初始化的类型: slice map channel
	// 初始化一个channel, 并设置缓冲区大小为1
	ch3 := make(chan int, 1)
	ch3 <- 10
	v := <-ch3
	fmt.Println(v)
	// 关闭通道
	close(ch3)
	// 关闭的通道, 可以继续取值
	// 如果所有值取完, 会返回对应类型的零值
	v1 := <-ch3
	fmt.Println(v1)
	fmt.Println("-------------------------------------------")
	// 使用goroutine和channel实现一个简易的生产者消费者模型
	// 生产者: 生成随机数
	// 消费者: 计算每个随机数的每个数字之和
	go producer()
	startWorker(20)
	printResult()
}
