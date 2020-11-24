package main

import "fmt"

type Employee struct {
	name, currency string
	salary float64
}

func main() {
	emp1 := Employee{"张三", "人民币", 2000}
	emp1.printSalary()
	printSalary(&emp1)
}

func (e Employee) printSalary()  {
	fmt.Println("员工：", e.name, "货币：", e.currency, "薪资：", e.salary)
}

func printSalary(e *Employee) {
	fmt.Println("员工：", e.name, "货币：", e.currency, "薪资：", e.salary)
}
