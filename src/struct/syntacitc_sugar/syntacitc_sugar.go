package main

import "fmt"

type Teacher struct {
	name string
	age int8
	sex byte
}
func main()  {
	// 语法糖
	t := new(Teacher)
	fmt.Println(t)
	fmt.Printf("t: %T %v %p\n", t,t,t)
	(*t).name = "wanglu"
	(*t).age = 12
	(*t).sex = 1
	fmt.Println(t)
	fmt.Printf("t: %T %v %p\n", t,t,t)
	t.name = "zhangsan"
	t.age = 23
	t.sex = 2
	fmt.Println(t)
	fmt.Printf("t: %T %v %p\n", t,t,t)
	syntacticSugar()
}

func syntacticSugar()  {
	arr := [4]int{10,30,50,60}
	arr2 := &arr
	fmt.Println((*arr2)[len(arr) - 1])
	fmt.Println(arr2[2])

	s := []int{100, 200,500, 700}
	s2 := &s
	fmt.Println((*s2)[len(s) - 1])

}
