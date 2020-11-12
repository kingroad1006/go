package main

import (
	"fmt"
	"net/http"
)

func main()  {
	// 绑定路由去触发方法
	http.HandleFunc("/index", indexHandle)
	// 绑定端口
	// 第一个参数为	监听地址，第二个参数表示服务端处理程序，通常为nil，这意味着服务端调用http.DefaultServeMux进行处理
	err := http.ListenAndServe("localhost:3013", nil)
	fmt.Println(err)
}

func indexHandle(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	w.Write([]byte(username))
	//fmt.Println("/index===========")
	//w.Write([]byte("这是默认首页"))
}
