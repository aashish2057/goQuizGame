package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "Enter csv file name: ")

	file, err := os.Open(*csvFileName)
	check(err)

	defer file.Close()

	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Printf("%+v\n", record)
	}

}
