package main

import "fmt"

func main() {
	arr := [5]int{10, 12, 15, 19, 24}
	arr_2d := [5][2]int{}
	for i := range arr_2d {
		arr_2d[i][0] = arr[i]
		arr_2d[i][1] = arr[i] * 2
	}
	for _, v := range arr_2d {
		fmt.Println(v)
	}
}
