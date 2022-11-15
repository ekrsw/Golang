package main

import (
	"fmt"
	"github.com/rocketlaunchr/dataframe-go"
	"gonum.org/v1/gonum/stat"
)

func main() {
	df := dataframe.NewSeriesFloat64("data", nil, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(df.Table())
}