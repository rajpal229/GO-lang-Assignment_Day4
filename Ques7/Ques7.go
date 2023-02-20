package main

import "fmt"

func power2(num int) int {
	var num2 = num * num
	return num2
}

func main() {
	inpSlice := make([]uint8, 3, 5)
	fmt.Println("Slice at the time of decalaration:", inpSlice)
	inpSlice[0] = 2
	inpSlice[1] = 4
	inpSlice[2] = 8
	fmt.Println("Slice after intial intialisation:", inpSlice)
	powerOf2Slice := make([]uint8, 0, 5)
	for _, value := range inpSlice {
		x := power2(int(value))
		powerOf2Slice = append(powerOf2Slice, uint8(x))
	}
	fmt.Println("PowerOf2Slice of intial inpSlice:", powerOf2Slice)
	inpSlice = append(inpSlice, 10)
	inpSlice = append(inpSlice, 16)
	fmt.Println("Slice after final intialisation:", inpSlice)
	for i := 3; i < len(inpSlice); i++ {
		x := power2(int(inpSlice[i]))
		powerOf2Slice = append(powerOf2Slice, uint8(x))
	}
	fmt.Println("PowerOf2Slice after final inpSlice:", powerOf2Slice)
}
