package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"_name"`
	Age int `json:"_age"`
	Sex uint `json:"-"` //不解析
	Address string //不改变key的标签
}

var user = User{
	Name:"Steven",
	Age: 35,
	Sex: 1,
	Address: "北京海淀区",
}

func main() {
	arr, _ := json.Marshal(user)
	fmt.Println(string(arr))
}
