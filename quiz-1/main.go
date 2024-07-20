package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type question struct {
	q   string
	ans string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getQuestions(f *os.File, questions *[]question) error {
	r := io.Reader(f)
	csvReader := csv.NewReader(r)
	// questions := make([]question, 10)
	for true {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		q := question{q: row[0], ans: row[1]}
		*questions = append(*questions, q)
	}

	return nil
}

func quiz(questions *[]question, totalQuestions *int, ansCount *int) (int, int) {
	for i, q := range *questions {
		// convert from string to number
		rowNum := i + 1
		ans, err := strconv.Atoi(q.ans)
		check(err)
		var userAnswer int
		fmt.Printf("Question %d:\t%s = ", rowNum, q.q)
		fmt.Scanln(&userAnswer)
		if userAnswer == ans {
			*ansCount += 1
		}
	}

	return *totalQuestions, *ansCount
}

func main() {
	exePath, _ := os.Executable()
	wd := filepath.Dir(exePath)
	var filename string
	var quizTime int = 30
	defaultFile := filepath.Join(wd + "/problems.csv")

	flag.StringVar(&filename, "f", defaultFile, "Specify filename")
	flag.StringVar(&filename, "file", defaultFile, "Specify filename (shorthand)")
	flag.IntVar(&quizTime, "t", quizTime, "Specify the quiz time (in seconds)")
	flag.IntVar(&quizTime, "time", quizTime, "Specify the quiz time (in seconds)")

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
	var questions []question
	err = getQuestions(f, &questions)
	check(err)

	fmt.Println("Quiz time: ", quizTime)
	fmt.Println("Press Enter to start the quiz")
	fmt.Scanln()

	ch := make(chan bool, 1)
	// var wg sync.WaitGroup
	// wg.Add(1)
	var totalQuestions int = 0
	var ansCount int = 0

	go func() {
		// defer wg.Done()
		quiz(&questions, &totalQuestions, &ansCount)

		// fmt.Println("------------------------------------------")
		// fmt.Printf("You answered %d correct out of %d question.\n", ansCount, totalQuestions)
		ch <- true
	}()

	// wg.Wait()

	select {
	case <-ch:
		fmt.Println("Done.")
	case <-time.After(time.Duration(quizTime * int(time.Second))):
		fmt.Println("\nTime out!")
	}

	fmt.Println("------------------------------------------")
	fmt.Printf("You answered %d correct out of %d question.\n", ansCount, len(questions))
}
