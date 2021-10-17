package main

/**
 * select 多路复用器
 *    	 select {
 * 		     case data := <- channel01:
 * 					...
 *      	 case data := <- channel02:
 * 					...
 * 			 ...
 *			 default:
 *            	  默认操作
 * 		 }
 */

import (
	"fmt"
	"time"
)

var ch1 = make(chan string, 100)
var ch2 = make(chan string, 100)

func f1(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("f1:%d", i)
		time.Sleep(time.Second)
	}
}

func f2(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("f2:%d", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go f1(ch1)
	go f2(ch2)
	for {
		select {
		case r := <-ch1:
			fmt.Println(r)
		case r := <-ch2:
			fmt.Println(r)
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}
