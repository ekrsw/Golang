package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatFloat(123.456123465798, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))
	fmt.Println(strconv.FormatFloat(123456789.123, 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64))
	fmt.Println(strconv.FormatFloat(123456789.123, 'G', 8, 64))
}