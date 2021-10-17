package main

/**
 * 反射:
 * 	  优势: 代码更加灵活
 *    劣势: 性能低
 */

import (
	"fmt"
	"reflect"
)

// 使用反射动态获取接口类型
func printType(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Printf("type: %v\n", t)
	fmt.Printf("name: %v, kind: %v\n", t.Name(), t.Kind())
	fmt.Println()
}

// 使用反射动态获取接口值信息
func printValue(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Printf("value: %v\n", v)
	fmt.Println()
}

// 使用反射修改值
func modifyValue(i interface{}) {
	v := reflect.ValueOf(i)
	kind := v.Kind()
	fmt.Println(kind)
	v.Elem().SetInt(10)
}

// 反射获取结构体具体信息
func printStruct(i interface{}) {
	t := reflect.TypeOf(i)
	for i := 0; i < t.NumField(); i++ {
		// 通过索引获取
		f := t.Field(i)
		fmt.Println(f)
		tag := f.Tag
		fmt.Printf("name: %s index: %d type: %v tag1: %v tag2: %v tag3: %v\n",
			f.Name, f.Index, f.Type, tag.Get("json"), tag.Get("tag"), tag.Get("v"),
		)
	}
	fmt.Println("----------------------------")
	// 通过字段名称获取字段信息
	if ageField, ok := t.FieldByName("Age"); ok {
		fmt.Println(ageField)
	}
}

type Person struct {
	Name string `json:"name" tag:"姓名" v:"001"`
	Age  uint8  `json:"age" tag:"年龄" v:"002"`
}

func main() {
	printType(100)
	printType(1.0)
	printType(false)
	printType("hello world")
	printType([2]int{1, 2})
	printType(map[string]int{})
	var p = Person{
		Name: "张三",
		Age:  18,
	}
	printType(p)
	printType(&p)
	fmt.Println("----------------------------")
	printValue(100)
	printValue(p)
	printValue(&p)
	fmt.Println("----------------------------")
	var a int64 = 100
	modifyValue(&a)
	fmt.Println(a)
	fmt.Println("----------------------------")
	printStruct(p)
}
