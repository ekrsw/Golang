package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.log")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags)
	log.Println("標準のログフォーマット")
}