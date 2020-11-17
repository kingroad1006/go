package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {
	TestContains()
	fmt.Println("-------------------")
	TestContainsAny()
	fmt.Println("-------------------")
	TestContainsRune()
	fmt.Println("-------------------")
	TestCount()
	fmt.Println("-------------------")
	TestHasPrefix()
	fmt.Println("-------------------")
	TestHasSuffix()
	fmt.Println("-------------------")
	TestIndex()
	fmt.Println("-------------------")
	TestIndexAny()
	fmt.Println("-------------------")
	TestIndexByte()
	fmt.Println("-------------------")
	TestIndexRune()
	fmt.Println("-------------------")
	TestIndexFunc()
	fmt.Println("-------------------")
	TestLastIndex()
	fmt.Println("-------------------")
	TestLastIndexAny()
	fmt.Println("-------------------")
	TestLastIndexByte()
	fmt.Println("-------------------")
	TestLastIndexFunc()
}

// 判断是否包含字符串
func TestContains() {
	fmt.Println(strings.Contains("Seafood", "foo"))
	fmt.Println(strings.Contains("Seafood", "bar"))
	fmt.Println(strings.Contains("Seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("Steven土2008", "土"))
}

// 判断字符串是否包含另一个字符串的人一个字符
func TestContainsAny()  {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "i & u"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

// 判断字符串是否包含Unicode码
func TestContainsRune() {
	fmt.Println(strings.ContainsRune("一丁", '丁'))
	fmt.Println(strings.ContainsRune("一丁", 19969))
	s := string("一丁")
	for _,v := range []rune(s) {
		fmt.Printf("%c\n", v)
	}
}

// 判断字符串包含另一个字符串的个数
func TestCount() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("one", ""))
}

// 判断字符串是否有前缀字符串
func TestHasPrefix() {
	fmt.Println(strings.HasPrefix("1000Phone new", "1000"))
	fmt.Println(strings.HasPrefix("1000Phone new", "1000a"))
}

// 判断一个字符串是否有后缀字符串
func TestHasSuffix()  {
	fmt.Println(strings.HasSuffix("1000Phone new", "new"))
	fmt.Println(strings.HasSuffix("1000Phone new", "news"))
}

// 返回一个字符串中另一个字符串首次出现的位置
func TestIndex() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr")) // 不存在返回-1
}

// 返回字符串中任一Unicode码值首次出现的位置
func TestIndexAny() {
	fmt.Println(strings.IndexAny("abcABC120", "教育基地A"))
}

// 返回字符串中字符首次出现的位置
func TestIndexByte() {
	fmt.Println(strings.IndexByte("123abc", 'a')) // 字符需要使用单引号
}

// 判断字符串中是否包含Unicode码值
func TestIndexRune() {
	fmt.Println(strings.IndexRune("abcABC120", 'C'))
	fmt.Println(strings.IndexRune("It教育培训", '教'))
	fmt.Println(strings.IndexRune("It教育培训", '学')) // 不存在返回-1
}

// 返回字符串中满足函数f(r) == true 字符首次出现的位置
func TestIndexFunc() {
	f := func (c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello123,中国", f))
}

// 返回字符串中最后一次子串出现的位置
func TestLastIndex()  {
	fmt.Println(strings.LastIndex("Steven learn english", "e"))
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
}

// 返回字符串中任一Unicode码值首次出现的位置
func TestLastIndexAny() {
	fmt.Println(strings.LastIndexAny("abcABC120", "教育基地A"))
}

// 返回字符串中字符最后一次出现的位置
func TestLastIndexByte() {
	fmt.Println(strings.LastIndexByte("abc123abc", 'a')) // 字符需要使用单引号
}

// 返回字符串中满足函数f(r) == true 字符最后一次出现的位置
func TestLastIndexFunc() {
	f := func (c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello123,中国", f))
}