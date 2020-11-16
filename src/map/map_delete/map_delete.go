package main

import "fmt"

func main() {
	m := map[string]string{
		"name" : "Zhangsan",
		"age" : "16",
		"sex" : "1",
	}
	fmt.Println("删除前：", m)
	if _,ok := m["age"]; ok {
		delete(m, "age")
	}
	fmt.Println("删除后：", m)
	m = map[string]string{}
	fmt.Println("清空后：", m)
}
