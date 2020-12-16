package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var sa []string
	sa := make([]string, 0, 0)
	printSliceMsg(sa)
	for i:=0;i<15;i++ {
		sa = append(sa, strconv.Itoa(i))
		printSliceMsg(sa)
	}
	printSliceMsg(sa)
}

func printSliceMsg(s []string)  {
	fmt.Printf("地址：%p 类型：%T 数值：%v 长度：%d 容量：%d\n", &s,s,s,len(s),cap(s))
}
