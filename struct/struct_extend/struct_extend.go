package main

import "fmt"

type Person struct {
	name string
	age int8
	sex byte
}

type Student struct {
	Person
	schoolName string
}

func main() {
	// 方式1
	p := Person{"张三", 25, 2}
	fmt.Println(p)
	fmt.Printf("p：%T %v %p\n", p, p, &p)
	s := Student{p, "十一中学"}
	fmt.Println(s)
	fmt.Printf("s: %T %v %p\n", s, s, &s)
	// 方式2
	s2 := Student{Person{"李思", 31, 1}, "中国人民大学"}
	fmt.Println(s2)
	fmt.Printf("s2: %T %v %p\n", s2, s2, &s2)
	// 方式3
	s3 := Student{
		Person : Person{
			name : "王五",
			age : 34,
			sex : 1,
		},
		schoolName : "人大附中",
	}
	fmt.Println(s3)
	fmt.Printf("s3: %T %v %p\n", s3, s3, &s3)
	// 方式4
	s4 := Student{}
	s4.name = "赵柳"
	s4.age = 24
	s4.sex = 1
	s4.schoolName = "衡水一中"
	fmt.Println(s4)
	fmt.Printf("s4: %T %v %p\n", s4, s4, &s4)
}
