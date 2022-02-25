package main

import (
	"fmt"
	"reflect"
)

func stringCompare(string1, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}
	map1 := map[string]int{}
	map2 := map[string]int{}
	for _, c := range string1 {
		map1[string(c)] += 1
	}
	for _, c := range string2 {
		map2[string(c)] += 1
	}

	return reflect.DeepEqual(map1, map2)
}

func main() {
	fmt.Println(stringCompare("abc", "cba"))
	fmt.Println(stringCompare("abc", "cbsa"))
	fmt.Println(stringCompare("abc", "cdm"))
}
