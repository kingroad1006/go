package main

import "fmt"

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {

	visit([]int{1, 3, 4, 6}, func(v int) {
		fmt.Println("hello", v)
	})

}
