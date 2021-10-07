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

	ticker := time.Tick(time.Second)

	for t := range ticker {
		fmt.Println("-------", t)
	}

}
