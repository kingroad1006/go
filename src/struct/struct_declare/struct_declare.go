package main

import "fmt"

type Teacher struct {
	name string
	age int8
	sex byte
}

func main() {
	// var 声明结构体
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1: %T %v %q\n", t1,t1,t1)
	t1.name = "张三"
	t1.age = 18
	t1.sex = 1
	fmt.Println(t1)

	//变量简短声明
	t2 := Teacher{}
	fmt.Println(t2)
	fmt.Printf("t2: %T %v %q\n", t2,t2,t2)
	t2.name = "张三"
	t2.age = 18
	t2.sex = 1
	fmt.Println(t2)

	//变量简短声明 并初始化
	t3 := Teacher{
		name: "李思",
		age : 12,
		sex : 2,
	}
	fmt.Println(t3)
	fmt.Printf("t2: %T %v %q\n", t3,t3,t3)

	// 简短声明不写属性名称
	t4 := Teacher{"wangwu", 23, 2}
	fmt.Println(t4)
	fmt.Printf("t2: %T %v %q\n", t4,t4,t4)
}
