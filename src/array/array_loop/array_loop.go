package main

import "fmt"

func main(){
	a := [4]float64{12.34, 43.33, 45.12, 98.23}
	b := [...]int{2,4,5}
	// 第一种遍历方式
	for i := 0;i < len(a);i++ {
		fmt.Println(a[i], "\t")
	}
	// 第二种遍历方式
	for _, value := range b {
		fmt.Println(value, "\t")
	}
}
