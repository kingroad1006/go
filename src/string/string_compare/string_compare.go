package main

import (
	"fmt"
	"strings"
)

func main() {
	TestCompare()
	TestEqualFold()
	TestRepeat()
	TestReplace()
	TestJoin()
}

// 按照字典顺序比较字符串a和b的大小
func TestCompare() {
	fmt.Println(strings.Compare("abc", "bcd"))
	fmt.Println("abc" < "bcd")
}

// 判断两个utf-8字符串是否相等，忽略大小写
func TestEqualFold() {
	fmt.Println(strings.EqualFold("GO", "go"))
}

// 将字符串重复count次返回
func TestRepeat() {
	fmt.Println("g" + strings.Repeat("o", 8) + "le")
}

// 将字符串s中的old字符串替换为new字符串，n<0时候替换所有的old字符串
func TestReplace() {
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", 2))
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", -1))
}

// 将a中的所有字符连接成一个字符串，使用字符串sep作为分隔符
func TestJoin() {
	a := []string{"abc", "ABC", "123"}
	fmt.Println(strings.Join(a, ","))
	fmt.Println(strings.Join(a, ""))
}
