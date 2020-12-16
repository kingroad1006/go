package main

import (
	"fmt"
	"strconv"
)

func main() {
	TestItoa()
	TestFormatInt()
	TestFormatUint()
	TestFormatFloat()
}

// int 转换成字符串
func TestItoa() {
	s := strconv.Itoa(123)
	fmt.Printf("%T %v\n",s,s)
}

// 返回给定基数的i的字符串表示
func TestFormatInt() {
	s := strconv.FormatInt(-19968,16)
	s = strconv.FormatInt(-40869,16)
	fmt.Printf("%T %v 长度：%d\n",s,s, len(s))
}

// 返回给定基数的i的字符串表示(无符号)
func TestFormatUint() {
	s := strconv.FormatUint(19968,16)
	s = strconv.FormatUint(40869,16)
	fmt.Printf("%T %v 长度：%d\n",s,s, len(s))
}

// 将float类型 转化为字符串
func TestFormatFloat() {
	s := strconv.FormatFloat(3.1415926, 'g', -1, 64)
	fmt.Printf("%T %v\n", s, s)
}
