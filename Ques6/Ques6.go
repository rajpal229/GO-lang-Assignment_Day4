package main

import (
	"fmt"
	"math/rand"
)

func check_odd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func main() {
	var arr [20]int
	for i := range arr {
		rn := rand.Intn(20-1) + 1
		arr[i] = rn
	}
	fmt.Println("Original arary:", arr)
	sliced_arr := arr[4:14]
	fmt.Println("Sliced Array:", sliced_arr)
	odd_slice := make([]int, 0)
	even_slice := make([]int, 0)
	for _, value := range sliced_arr {
		if check_odd(value) {
			odd_slice = append(odd_slice, value)
		} else {
			even_slice = append(even_slice, value)
		}
	}
	fmt.Println("Odd number slice:", odd_slice)
	fmt.Println("Even Number Slice:", even_slice)
}
