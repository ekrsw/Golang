package main

import "fmt"

func Double(x []int) {
	for i, v := range x {
		x[i] = v * 2
	}
}

func main() {
	sl := []int{1, 2, 3, 4}
	Double(sl)
	fmt.Println(sl)
}