package main

/**
 * 文件操作
 */

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filepath := "./day03/xx.txt"
	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("文件打开失败, err : ", err)
		return
	}
	// 关闭文件, defer会在return前执行, 类似Java中的finally
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// 读取文件
	// bufio 方式
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		fmt.Print(readString)
		//if err == io.EOF {
		//	return
		//}
		if err != nil {
			break
		}
	}
	fmt.Println("\n-----------------------")
	// ioutil方式
	readFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(readFile))

	// 写文件
	openFile, openFileErr := os.OpenFile("./day03/oo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if openFileErr != nil {
		fmt.Println(openFileErr)
		return
	}
	fd, err := openFile.WriteString("hello world!\n")
	fmt.Println(fd)
	if err != nil {
		return
	}

	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(openFile)

}
