package main

import "fmt"

func main() {
	switch n := 3; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know.")
	}
}
