package main

import "fmt"

func main() {
	a := [5][2]int{{1,2},{3,4},{5,6},{7,8},{9,10}}
	fmt.Println(len(a))
	fmt.Println(len(a[0]))
	for i := 0; i< len(a); i++ {
		for j:=0; j< len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
