package main

import (
	"fmt"
	"strings"
)

func main() {
	TestTitle()
	TestToTitle()
	TestToLower()
	TestToUpper()
}

// 将字符串每个单词的首字符大写返回
func TestTitle() {
	fmt.Println(strings.Title("hello world"))
}

// 将字符串转化大写返回
func TestToTitle() {
	fmt.Println(strings.ToTitle("hello world"))
}

// 将字符串转化成小写
func TestToLower() {
	fmt.Println(strings.ToLower("HELLO WORLD"))
}

// 将字符串转化大写返回
func TestToUpper()  {
	fmt.Println(strings.ToUpper("hello world"))
}