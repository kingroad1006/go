package main

import "fmt"

type Human struct{
	name string
	age int8
	sex byte
}

func main() {
	// 结构体值传递
	h1 := Human{"zhangsan", 12, 1}
	fmt.Println(h1)
	h2 := h1
	h2.name = "lisi"
	h2.age = 18
	h2.sex = 2
	fmt.Println(h2)
	changeName(h2)
}

func changeName(h Human) {
	h.name = "wangwu"
	fmt.Println(h)
}
