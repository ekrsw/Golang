package main

import (
	"Golang/alib"
	"fmt"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))

	sl := []int{1, 2, 3, 4, 6}
	fmt.Println(alib.Average(sl))
}
