package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestFields()
	TestFieldsFunc()
	TestSplit()
	TestSplitN()
	TestSplitAfter()
	TestSplitAfterN()
}

// 将字符串用空白符分隔，并返回一个切片
func TestFields() {
	fmt.Println(strings.Fields("zhangsan lisi wanglu mazi"))
}

// 将字符串以满足f(r) == true的字符串分隔，返回一个切片
func TestFieldsFunc()  {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("abc@123*ABC&xyz%XYZ", f))
}

// 将字符串以sep作为分隔符分隔，分割后字符最后去掉sep
func TestSplit() {
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

// 将字符串s以sep作为分隔符进行分隔，分隔后的字符串最后附上sep, n决定返回的切片数
func TestSplitN() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 1))
}

// 将字符串s以sep作为分隔符进行分隔，分隔后的字符串最后附上sep
func TestSplitAfter() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
}

// 将字符串s以sep作为分隔符进行分隔，分隔后的字符串最后附上sep, n决定返回的切片数
func TestSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 1))
}

