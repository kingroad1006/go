package main

import (
	"fmt"
	"net/http"
)

func main() {
	testClientGet()
}

func testClientGet() {
	client := http.Client{}

	response, err := client.Get("https://www.baidu.com")

	CheckErr(err)
	fmt.Printf("响应状态码：%v\n", response.StatusCode)
	if response.StatusCode == 200 {
		fmt.Println("网络请求成功")
		defer response.Body.Close()
	}
}

func CheckErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现错误：", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
