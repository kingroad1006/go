package main

import (
	"encoding/json"
	"fmt"
)

type DebugInfo struct {
	Level string
	Msg string
	author string // 首字母小写不会被json解析
}
func main() {
	dbginfs := []DebugInfo{
		{"debug", `File: "test.txt" Not found`, "Cynhard"},
		{"", "Logic error", "Gopher"},
	}

	if data, err := json.Marshal(dbginfs); err == nil {
		fmt.Printf("%s \n", data)
	}
}
