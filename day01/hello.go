package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world!")
	// 字符串
	s1 := "abcdefg"
	// 多行字符串
	s2 := `a
	b
	c
	d
	`
	// 字符串长度
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	// 字符串拼接
	s3 := "xx"
	s4 := "oo"
	fmt.Println(s3 + s4)
	s5 := fmt.Sprintf("%s----%s", s3, s4)
	fmt.Println(s5)

	// 分割, 返回切片
	s6 := strings.Split("a,b,c,d", ",")
	fmt.Println(s6)

	// 包含
	s7 := strings.Contains(s5, "a")
	fmt.Println(s7)

	// 判断前缀和后缀
	s8 := strings.HasPrefix("http://baidu.com", "http://")
	fmt.Printf("s8: %v\n", s8)
	s9 := strings.HasSuffix("http://baidu.com", "com")
	fmt.Printf("s9: %v\n", s9)

	// 字符串出现的位置
	s10 := strings.LastIndex("xxooxxoo", "oo")
	fmt.Printf("s10: %v\n", s10)

	// join
	s11 := []string{"Java", "JavaScript", "Golang", "Python"}
	fmt.Printf("s11 join: %v\n", strings.Join(s11, "、"))

}
