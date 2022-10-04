package main

import (
	"golang.org/x/text/encoding/japanese"
    "golang.org/x/text/transform"

	"encoding/csv"
	"io"
	"log"
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

	records, _ := reader.ReadAll()
	for _, row := range records {
		fmt.Println(row)
	}
}