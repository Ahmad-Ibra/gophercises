package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
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
			panic(err)
		}

		prob := Problem{q: record[0], a: record[1]}
		problems = append(problems, prob)
	}
	return problems
}

func main() {
	// TODO: add flag for custom time limit
	file, err := os.ReadFile("problems.csv")
	if err != nil {
		panic(err)
	}

	problems := parseFile(string(file))

	var score int
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		var userAnswer string
		_, err := fmt.Scanln(&userAnswer)
		if err != nil {
			return
		}

		if userAnswer == problem.a {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}