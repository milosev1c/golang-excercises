package main

import (
	"fmt"
)

const SZ = 5

func reverse(arr *[SZ]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5}

	reverse(&arr)

	fmt.Printf("%v", arr)

}
