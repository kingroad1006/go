package main

import "fmt"

func main() {
	// 匿名函数
	f := func(a,b int) (s int) {
		return a * b
	}(2, 3)
	fmt.Println(f)
	// 匿名结构体
	s := struct {
		name string
		age int8
	}{name: "张三", age : 15}
	fmt.Println(s)
}
