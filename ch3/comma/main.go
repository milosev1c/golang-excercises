package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	var decPart string
	c := 0
	if decPos := strings.LastIndex(s, "."); decPos != -1 {
		decPart = s[decPos:]
		s = s[:decPos]
	}
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		c++
		if c%3 == 0 && i != 0 {
			s = s[:i] + "," + s[i:]
		}
	}
	return s + decPart
}

func main() {
	fmt.Println(comma("123456"))
	fmt.Println(comma("-3330"))
}
