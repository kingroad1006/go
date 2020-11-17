package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestTrim()
	TestTrimFunc()
	TestTrimLeft()
	TestTrimLeftFunc()
	TestTrimRight()
	TestTrimRightFunc()
	TestTrimSpace()
	TestTrimPrefix()
	TestTrimSuffix()
}

// 将字符串中 首尾包含在 cutset 中的任一字符去掉返回
func TestTrim() {
	fmt.Println(strings.Trim(" Steven Wang ", " "))
}

// 将字符串中首尾满足 f(r) == true的字符去掉返回
func TestTrimFunc()  {
	f := func(c rune) bool{
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimFunc("$%^& Steven Wang $%^&*", f))
}

// 将字符串中左边包含在 cutset 中的任一字符去掉返回
func TestTrimLeft() {
	fmt.Println(strings.TrimLeft(" Steven Wang ", " "))
}

// 将字符串中左边满足 f(r) == true的字符去掉返回
func TestTrimLeftFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimLeftFunc("$%^& Steven Wang $%^&*", f))
}

// 将字符串中右边包含在 cutset 中的任一字符去掉返回
func TestTrimRight() {
	fmt.Println(strings.TrimRight(" Steven Wang ", " "))
}

// 将字符串中右边满足 f(r) == true的字符去掉返回
func TestTrimRightFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimRightFunc("$%^& Steven Wang $%^&*", f))
}

// 将字符串首尾空白去掉返回
func TestTrimSpace() {
	fmt.Println(strings.TrimSpace(" Steven Wang "))
}

// 将字符串中的前缀字符串prefix去掉
func TestTrimPrefix() {
	fmt.Println(strings.TrimPrefix("hello world", "hello"))
}

// 将字符串中的后缀字符串suffix去掉
func TestTrimSuffix() {
	fmt.Println(strings.TrimSuffix("hello world", "world"))
}
