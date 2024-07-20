package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Define flags
	exePath, _ := os.Executable()
	wd := filepath.Dir(exePath)
	var filename string
	defaultFile := filepath.Join(wd + "/problems.csv")
	flag.StringVar(&filename, "f", defaultFile, "Specify filename")
	flag.StringVar(&filename, "file", defaultFile, "Specify filename (shorthand)")

	// Parse flags
	flag.Parse()

	// Check if a filename was provided
	if filename == "" {
		fmt.Println("No filename provided, using default: problems.csv")
	} else {
		fmt.Println("Using filename:", filename)
	}

	f, err := os.Open(filename)
	check(err)

	r := io.Reader(f)
	csvReader := csv.NewReader(r)
	// for i, q := range csvR {
	// 	fmt.Printf("row %d:\t %s\n", i, q)
	// }
	// row, err := csvReader.Read()
	check(err)

	var rowNum = 0
	var correctAnsCount = 0
	for true {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		rowNum += 1

		// fmt.Printf("row %d:\t%s\t%s\n", rowNum, row[0], row[1])
		// col2 := strings.(row[1])
		// convert from string to number
		var question, ansStr = row[0], row[1]
		ans, err := strconv.Atoi(ansStr)
		check(err)
		var userAnswer int
		fmt.Printf("Question %d:\t%s = ", rowNum, question)
		fmt.Scanln(&userAnswer)
		if userAnswer == ans {
			correctAnsCount += 1
		}
	}

	fmt.Println("------------------------------------------")
	fmt.Printf("You answerd %d correct out of %d question.\n", correctAnsCount, rowNum)
	// fmt.Println("EOF")
}
