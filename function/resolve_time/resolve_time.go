package main

import "fmt"

const (
	Minute = 60
	Hour   = Minute * 60
	Day    = Hour * 12
)

func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / Day
	hour = seconds / Hour
	minute = seconds / Minute

	return
}

func main() {
	fmt.Println(Minute)
	fmt.Println(resolveTime(300000))
}
