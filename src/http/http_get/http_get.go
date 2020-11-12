package main

import (
	"fmt"
	"net/http"
)

func main() {
	testHttpGet()
}

func testHttpGet(){
	response, err := http.Get("https://www.baidu.com")
	CheckErr(err)
	fmt.Printf("响应状态码：%v\n", response.StatusCode)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		fmt.Println("网络请求成功")
		CheckErr(err)
	} else {
		fmt.Println("网络请求失败", response.Status)
	}

}

func CheckErr(err error)  {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常", ins.Error())
		}
	}()

	if err != nil {
		panic(err)
	}
}
