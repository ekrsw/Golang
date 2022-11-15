package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 100
	s := strconv.Itoa(i)

	s2 := strconv.FormatInt(i)

	fmt.Println(s, 10)
}