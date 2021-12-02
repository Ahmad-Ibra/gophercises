package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Problem struct {
	q string // question
	a string // answer
}

func parseFile(file string) []Problem {
	r := csv.NewReader(strings.NewReader(file))
	var problems []Problem

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Failed to read the CSV file.")
			os.Exit(1)
		}

		prob := Problem{q: record[0], a: record[1]}
		problems = append(problems, prob)
	}
	return problems
}

func runQuiz(problems []Problem, score *int, c chan<- bool) {
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		var userAnswer string
		_, err := fmt.Scanln(&userAnswer)
		if err != nil {
			fmt.Printf("Failed to scan input answer.")
			os.Exit(1)
		}

		if userAnswer == problem.a {
			*score++
		}
	}
	c <- true
}

func main() {
	pLimit := flag.Int("limit", 9999, "the time limit in seconds")
	flag.Parse()

	file, err := os.ReadFile("problems.csv")
	if err != nil {
		fmt.Printf("Failed to open the CSV file.")
		os.Exit(1)
	}

	problems := parseFile(string(file))
	var score int
	c := make(chan bool)
	go runQuiz(problems, &score, c)

	select {
	case <-c:
		break
	case <-time.After(time.Duration(*pLimit) * time.Second):
		break
	}

	fmt.Printf("\nYou scored %d out of %d.\n", score, len(problems))
}