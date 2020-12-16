package main

import "fmt"

type Flower struct {
	name, color string
}

func main() {
	// 结构体作为参数
	f1 := Flower{"rose", "red"}
	fmt.Println(f1)
	fmt.Printf("f1: %T %v %p\n", f1, f1, &f1)
	// 结构体对象作为参数
	changeInfo1(f1)
	fmt.Printf("f1: %T %v %p\n", f1, f1, &f1)
	// 指针作为参数
	changeInfo2(&f1)
	fmt.Printf("f1: %T %v %p\n", f1, f1, &f1)
	fmt.Println("---------------------------")
	// 结构体作为返回值
	// 结构体对象返回
	f2 := getFlower1()
	f3 := getFlower1()
	fmt.Println("更改前：", f2, f3)
	fmt.Printf("f2: %T %v %p\n", f2, f2, &f2)
	fmt.Printf("f2: %T %v %p\n", f3, f3, &f3)
	f2.name = "杜鹃"
	fmt.Println("更改后：", f2, f3)
	fmt.Printf("f2: %T %v %p\n", f2, f2, &f2)
	fmt.Printf("f2: %T %v %p\n", f3, f3, &f3)
	fmt.Println("---------------------------")
	// 结构体指针返回
	f4 := getFlower2()
	f5 := getFlower2()
	fmt.Println("更改前：", f4, f5)
	fmt.Printf("f2: %T %v %p\n", f4, f4, &f4)
	fmt.Printf("f2: %T %v %p\n", f5, f5, &f5)
	f4.name = "芙蓉"
	fmt.Println("更改前：", f4, f5)
	fmt.Printf("f2: %T %v %p\n", f4, f4, &f4)
	fmt.Printf("f2: %T %v %p\n", f5, f5, &f5)
}

func getFlower1() (f Flower) {
	f = Flower{"牡丹", "白"}
	fmt.Printf("函数getFlower1内: %T %v %p\n", f, f, &f)
	return
}

func getFlower2() (f *Flower) {
	f = &Flower{"月季", "粉"}
	fmt.Printf("函数getFlower2内: %T %v %p\n", f, f, &f)
	return
}

func changeInfo1(f Flower) {
	f.name = "玫瑰"
	f.color = "红色"
	fmt.Println(f)
	fmt.Printf("f: %T %v %p\n", f, f, &f)
}

func changeInfo2(f *Flower) {
	f.name = "太阳花"
	f.color = "白色"
	fmt.Println(f)
	fmt.Printf("f: %T %v %p\n", f, f, &f)
}