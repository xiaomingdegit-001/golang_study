package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%#v\n", now)

	fmt.Println(now.Unix())

	fmt.Println(now.Second())

	// 日期格式化 [2006-01-02 15:04:05.000] (格式有点奇葩, 必须为go语言诞生时间)
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02"))

	// 定时任务
	ticker := time.Tick(time.Second)

	for t := range ticker {
		fmt.Println("-------", t)
	}

}
