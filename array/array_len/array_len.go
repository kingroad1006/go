package main

import "fmt"

func main() {
	a := [4]float64{12.34, 43.33, 45.12, 98.23}
	b := [...]int{2, 4, 5}
	fmt.Printf("数组a的长度：%d\n数组b的长度：%d\n", len(a), len(b))
}
