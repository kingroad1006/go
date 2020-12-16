package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 0, 20)
	printSlice("numbers:", numbers)
	numbers = append(numbers, 0)
	printSlice("numbers:", numbers)
	numbers = append(numbers, 1)
	printSlice("numbers:", numbers)
	numbers = append(numbers, 2,3,4,5)
	printSlice("numbers:", numbers)
	s1 := []int{6,7,8,9}
	numbers = append(numbers, s1...)
	printSlice("numbers:", numbers)
	numbers = numbers[1:]
	printSlice("numbers:", numbers)
	numbers = numbers[:len(numbers)-1]
	printSlice("numbers:", numbers)
	a := int(len(numbers)/2)
	fmt.Println("中间数：", a)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlice("numbers:", numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers)) * 2)
	count := copy(numbers1, numbers)
	fmt.Println("复制个数：", count)
	printSlice("numbers1:", numbers1)
	numbers[len(numbers)-1] = 99
	numbers1[0] = 100
	printSlice("numbers:", numbers)
	printSlice("numbers1:", numbers1)
}

func printSlice(name string, s []int)  {
	fmt.Println(name, "\t")
	fmt.Printf("地址：%p 类型：%T 数值：%v 长度：%d\n", &s,s,s,len(s))
}
