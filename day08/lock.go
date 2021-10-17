package main

/**
 * 锁
 */

import (
	"fmt"
	"sync"
	"time"
)

var num int64
var num1 int64
var m map[string]int
var wg sync.WaitGroup

// 定义一个互斥锁
var lock sync.Mutex

// 定义一个读写互斥锁
var rwLock sync.RWMutex

// sync.Once 并发安全
// 功能有点类似Java中实现线程安全的单例模式
var loadOnce sync.Once

func add() {
	for i := 0; i < 10000; i++ {
		num++
	}
	wg.Done()
}

func add1() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		num1++
		lock.Unlock()
	}
	wg.Done()
}

func readByLock() {
	defer wg.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func writeByLock() {
	defer wg.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 50)
	lock.Unlock()
}

func readByRWLock() {
	defer wg.Done()
	rwLock.RLock()
	time.Sleep(time.Millisecond * 1)
	rwLock.RUnlock()
}

func writeByRWLock() {
	defer wg.Done()
	rwLock.Lock()
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
}

func load() {
	if m != nil {
		m = make(map[string]int)
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(num)
	fmt.Println("------------------------")
	wg.Add(2)
	go add1()
	go add1()
	wg.Wait()
	fmt.Println(num1)
	fmt.Println("------------------------")
	// 读操作远大于写操作的时候使用读写锁, 可以提升性能
	start := time.Now()
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go readByLock()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeByLock()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("互斥锁耗时: %v\n", end.Sub(start))
	fmt.Println("------------------------")
	start1 := time.Now()
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go readByRWLock()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeByRWLock()
	}
	wg.Wait()
	end1 := time.Now()
	fmt.Printf("读写互斥锁耗时: %v\n", end1.Sub(start1))
	fmt.Println("------------------------")
	loadOnce.Do(load)
	fmt.Println("------------------------")

	// 内置的 map 是线程不安全的
	// 并发场景下建议使用 sync.Map

	// atomic包
	// 和Java中Atomic类功能一样

}
