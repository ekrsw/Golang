package main

import (
	"log"
	"os"
)

func main () {
	file, err := os.Open("test1.log")
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Println(err)
	}
	defer file.Close()
}