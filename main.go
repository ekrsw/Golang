package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i, i)
}