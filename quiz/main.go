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

type problem struct {
	q string // question
	a string // answer
}

func parseFile(file string) []problem {
	r := csv.NewReader(strings.NewReader(file))
	var problems []problem

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			exit(fmt.Sprintf("Failed to read the CSV file."))
		}

		prob := problem{q: record[0], a: record[1]}
		problems = append(problems, prob)
	}
	return problems
}

func runQuiz(problems []problem, score *int, c chan<- bool) {
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		var userAnswer string
		_, err := fmt.Scanln(&userAnswer)
		if err != nil {
			exit(fmt.Sprintf("Failed to scan input answer."))
		}

		if userAnswer == problem.a {
			*score++
		}
	}
	c <- true
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	pLimit := flag.Int("limit", 9999, "the time limit in seconds")
	flag.Parse()

	file, err := os.ReadFile("problems.csv")
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file."))
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