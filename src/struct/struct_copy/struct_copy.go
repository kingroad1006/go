package main

import "fmt"

type Dog struct {
	name string
	color string
	age int8
	kind string
}

func main() {
	// 深拷贝
	d1 := Dog{"豆豆", "黑色",2,"二哈"}
	fmt.Println(d1)
	fmt.Printf("d1：%T %v %p\n", d1,d1,&d1)
	d2 := d1
	d2.name = "毛毛"
	fmt.Println(d2)
	fmt.Printf("d2：%T %v %p\n", d2,d2,&d2)

	// 浅拷贝
	d3 := &d1
	d3.age = 5
	fmt.Println(d1)
	fmt.Printf("d1：%T %v %p\n", d1,d1,&d1)
	fmt.Println(d3)
	fmt.Printf("d3：%T %v %p\n", d3,d3,d3)

	// 浅拷贝 new
	d4 := new(Dog)
	d4.name = "多多"
	d4.color = "棕色"
	d4.age = 4
	d4.kind = "泰迪"
	fmt.Println(d4)
	fmt.Printf("d4：%T %v %p\n", d4,d4,d4)
	d5 := d4
	fmt.Println(d5)
	fmt.Printf("d5：%T %v %p\n", d5,d5,d5)

}