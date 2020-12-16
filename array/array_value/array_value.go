package main

import "fmt"

func main() {
	a := [...]string{"aa","bb","cc","dd"}
	b := a
	b[0] = "ee"
	fmt.Println(a)
	fmt.Println(b)
}
