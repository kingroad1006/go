package main

import (
	"fmt"
)

func main() {
	// 声明同时初始化
	var country = map[string]string{
		"China" : "Beijing",
		"China1" : "Beijing1",
		"China2" : "Beijing2",
		"China3" : "Beijing3",
		"China4" : "Beijing4",
	}
	fmt.Println(country)
	// 短变量声明初始化
	rating := map[string]float64{
		"c" : 5,
		"go" : 4.5,
		"python" : 3.9,
		"c++" : 4.8,
	}
	fmt.Println(rating)
	// 创建后在赋值
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["China1"] = "Beijing1"
	fmt.Println(countryMap)
	// 遍历
	for k, v := range countryMap {
		fmt.Println("国家", k, "首都", v)
	}
	// 遍历 只展示value
	for _, v := range rating {
		fmt.Println("rating", v)
	}
	// 遍历 只展示key
	for k := range rating {
		fmt.Println("language", k)
	}

	//value, ok := countryMap["England"]
	value, ok := countryMap["China"]
	fmt.Printf("%q\n", value)
	fmt.Printf("%T, %v \n", ok, ok)
	if ok {
		fmt.Println("首都", value)
	} else {
		fmt.Println("首都信息未检索到")
	}

	if value, ok := countryMap["England"]; ok {
		fmt.Println("首都", value)
	} else {
		fmt.Println("首都信息未检索到")
	}

}
