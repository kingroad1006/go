package main

import (
	"fmt"
	"strings"
)

func stringProccess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str

		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	list := []string{
		"go aaaa",
		"go bbbb",
		"go cccc",
		"go dddd",
		"go eeee",
		"go ffff",
		"go gggg",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	stringProccess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}
