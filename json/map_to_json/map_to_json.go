package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	// 定义map并初始化
	m := map[string][]string{
		"level":{"debug"},
		"massage":{"File not found", "Stack overflow"},
	}
	// 将map解析成json格式
	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%s\n", data)
	}

	// 将map解析成json格式(格式化)
	if data, err := json.MarshalIndent(m, "", " "); err == nil {
		fmt.Printf("%s\n", data)
	}
}
