package main

import (
	"fmt"
	"flag"
)

func main() {
	var (
		max int
		msg string
		x bool
	)

	flag.IntVar(&max, "n", 64, "max max max")
	flag.StringVar(&msg, "m", "", "msg msg msg")
	flag.BoolVar(&x, "x", false, "x x x")

	flag.Parse()

	fmt.Println(max)
	fmt.Println(msg)
	fmt.Println(x)
}