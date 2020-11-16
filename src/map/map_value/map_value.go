package main

import "fmt"

func main() {
	m := map[string]string{
		"name" : "Zhangsan",
		"age" : "16",
		"sex" : "1",
	}
	fmt.Println(m)
	m1 := m
	m1["name"] = "Lisi"
	fmt.Println(m1)
	fmt.Println(m)
}
