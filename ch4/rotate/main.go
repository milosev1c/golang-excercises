package main

import "fmt"

func rotate(arr []int, turns int) []int {
	for i := 0; i < turns; i++ {
		ln := len(arr)
		arr = append([]int{0}, arr...)
		arr[0], arr[ln] = arr[ln], arr[0]
		arr = arr[:ln]
	}

	return arr

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = rotate(arr, 5)
	fmt.Printf("%v", arr)
}
