package main

import "fmt"

func Sum(s ..int) int {
	n := 0
	for _, u := range s {
		n += u
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
}