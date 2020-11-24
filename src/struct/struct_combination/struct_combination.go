package main

import "fmt"

type Address struct {
	province,city string
}

type Person struct {
	name string
	age int8
	address *Address
}

func main() {
	// 模拟结构体之间的聚合关系
	p := Person{}
	p.name = "张三"
	p.age = 34
	// 赋值方式1
	addr := Address{}
	addr.province = "北京市"
	addr.city = "海淀区"
	p.address = &addr
	fmt.Println(p)
	fmt.Println("name",p.name,"age", p.age,"province", p.address.province, "city", p.address.city)
	p.address.city = "昌平区"
	fmt.Println("name",p.name,"age", p.age,"province", p.address.province, "city", p.address.city, "city2", addr.city)
	addr.city = "丰台区"
	fmt.Println("name",p.name,"age", p.age,"province", p.address.province, "city", p.address.city, "city2", addr.city)

	// 赋值方式2
	p.address = &Address{
		province: "上海市",
		city: "浦东区",
	}
	fmt.Println("name",p.name,"age", p.age,"province", p.address.province, "city", p.address.city, "city2", addr.city)

}
