package main

import (
	"fmt"
	"reflect"
	"strings"
)

func max_secmax(arr []string) (string, string) {
	var (
		max string = arr[0]
		sec string = arr[0]
	)
	for i := 0; i < len(arr); i++ {
		if arr[i] > sec {
			sec = arr[i]
			if arr[i] > max {
				sec = max
				max = arr[i]
			}
		}
	}
	return max, sec
}

func min_secmin(arr []string) (string, string) {
	var (
		min string = arr[0]
		sec string = arr[0]
	)
	for i := 0; i < len(arr); i++ {
		if arr[i] < sec {
			sec = arr[i]
			if arr[i] < min {
				sec = min
				min = arr[i]
			}
		}
	}
	return min, sec
}

func main() {
	var s string
	fmt.Println("Enter comma seperated numbers: ")
	fmt.Scan(&s)
	str_arr := strings.Split(s, ",")
	fmt.Printf("This %v is Type - %v\n", str_arr, reflect.TypeOf(str_arr).Kind())
	fmt.Println("")
	max, secmax := max_secmax(str_arr)
	fmt.Println("Maximum:", max, "Next Highest:", secmax)
	min, secmin := min_secmin(str_arr)
	fmt.Println("Minimum:", min, "Next Least:", secmin)
	// str_slice := str_arr[:]
	// fmt.Printf("This %v is Type - %v\n", str_slice, reflect.TypeOf(str_slice).Kind())
}
