package main

import "fmt"

func main() {
	//var numbers = make([]int, 3, 5)
	//numbers := make([]int, 3, 5)
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Printf("%T\n", numbers)
	fmt.Printf("len: %d\ncap: %d\nslice: %v\n", len(numbers), cap(numbers), numbers)
	fmt.Println("numbers[1:4]==", numbers[1:4])
	fmt.Println("numbers[:3]==", numbers[:3])
	fmt.Println("numbers[4:]==", numbers[4:])
	numbers2 := numbers[:2]
	fmt.Printf("len: %d\ncap: %d\nslice: %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers3 := numbers[2:5]
	fmt.Printf("len: %d\ncap: %d\nslice: %v\n", len(numbers3), cap(numbers3), numbers3)
}
