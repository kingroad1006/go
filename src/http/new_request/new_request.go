package main

import (
	"fmt"
	"net/http"
)

func main() {
	testHttpNewRequest()
}

func testHttpNewRequest() {
	// 1创建一个客户端
	client := http.Client{}
	// 2创建一个请求 请求方式可以是GET 也可以是POST
	request, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	CheckErr(err)
	// 3客户端发送请求
	cookName := &http.Cookie{Name: "username", Value: "Steven"}
	// 添加cookie
	request.AddCookie(cookName)
	response, err := client.Do(request)
	CheckErr(err)
	// 设置请求头
	request.Header.Set("Accept-Language", "zh-cn")
	defer response.Body.Close()
	// 查看请求头数据
	fmt.Printf("Header:%v\n", request.Header)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)
	// 4操作数据
	if response.StatusCode == 200 {
		fmt.Println("网络请求成功")
	} else {
		fmt.Println("网络请求失败", response.Status)
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
