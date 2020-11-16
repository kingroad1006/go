package main

import "fmt"

func main() {
	a := [4]int{1,2,3,4}
	b := []int{5,6,7,8}
	fmt.Printf("a-地址：%p 类型：%T 数值：%v 长度：%d\n", &a,a,a,len(a))
	fmt.Printf("b-地址：%p 类型：%T 数值：%v 长度：%d\n", &b,b,b,len(b))
	c := a
	d := b
	fmt.Printf("c-地址：%p 类型：%T 数值：%v 长度：%d\n", &c,c,c,len(c))
	fmt.Printf("d-地址：%p 类型：%T 数值：%v 长度：%d\n", &d,d,d,len(d))
	c[0] = 100
	d[0] = 200
	fmt.Println("a=",a,"c=",c)
	fmt.Println("b=",b,"d=",d)

}
