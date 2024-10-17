package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	reader, _ := os.Open("local_data/data.csv")
	csvReader := csv.NewReader(reader)
	rows, _ := csvReader.ReadAll()
	for _, row := range rows {
		fmt.Println(row)
	}
	fmt.Println("hello")
}
