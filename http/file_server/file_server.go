package main

import "net/http"

func main()  {
	testFileService()
}

func testFileService()  {
	//如果路径中有index.html文件，会优先显示html邮件，否则会看到文件目录
	http.ListenAndServe(":2003", http.FileServer(http.Dir("../")))
}
