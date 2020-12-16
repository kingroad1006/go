package main

import (
	"fmt"
	"strconv"
)

func main() {
	TestAtoi()
	TestParseInt()
	TestParseUint()
	TestParseFloat()
	TestParseBool()
}

// 将字符串转化成int类型
func TestAtoi()  {
	a,_ := strconv.Atoi("100")
	fmt.Printf("%T %v\n", a, a+2)
}

// 解释给定基数（2到36）的字符串s并返回相应的值i
func TestParseInt() {
	num,_ := strconv.ParseInt("-4e00", 16, 64)
	fmt.Printf("%T %v\n", num, num)
	num,_ = strconv.ParseInt("01100001", 2, 64)
	fmt.Printf("%T %v\n", num, num)
	num,_ = strconv.ParseInt("-01100001", 10, 64)
	fmt.Printf("%T %v\n", num, num)
	num,_ = strconv.ParseInt("4e00", 10, 64)
	fmt.Printf("%T %v\n", num, num)
}

// 类似ParseInt 无符号
func TestParseUint() {
	num,_ := strconv.ParseUint("-4e00", 16, 64)
	fmt.Printf("%T %v\n", num, num)
	num,_ = strconv.ParseUint("01100001", 2, 64)
	fmt.Printf("%T %v\n", num, num)
	num,_ = strconv.ParseUint("-01100001", 10, 64)
	fmt.Printf("%T %v\n", num, num)
	num,_ = strconv.ParseUint("4e00", 10, 64)
	fmt.Printf("%T %v\n", num, num)
}

// 将字符串 s转化为float类型
func TestParseFloat() {
	pi := "3.1415926"
	num,_ := strconv.ParseFloat(pi, 64)
	fmt.Printf("%T %v\n", num, num)
}

// 将字符串转换成bool类型
func TestParseBool() {
	//flag,_ := strconv.ParseBool("Steven")
	flag,_ := strconv.ParseBool("true")
	fmt.Printf("%T %v\n", flag, flag)
}