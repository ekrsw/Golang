package main

import (
	"os"
	"log"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile)
	_, err := os.Open("sdfasdf")
	if err != nil {
		logger.Fatalln("Exit", err)
	}
}