package main

import (
	"fmt"
	"strings"
)

func remove_duplicates(str []string) string {

	for i := 0; i < (len(str) - 1); {
		if str[i] == str[i+1] {
			copy(str[i:], str[i+1:])
			str = str[:len(str)-1]
			continue
		}
		i++
	}
	return strings.Join(str, "")
}

func main() {
	str := "aaadddbbbeee"
	res := remove_duplicates(strings.Split(str, ""))
	fmt.Printf("%v", res)
}
