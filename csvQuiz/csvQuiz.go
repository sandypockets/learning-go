package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func parseCsvLines(lines [][]string) []problem {
	problemSlice := make([]problem, len(lines))
	for i, line := range lines {
		problemSlice[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problemSlice
}

func errMsg(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		errMsg(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFileName))
	}

	fmt.Printf("Successfully opened the CSV file: %s\n", *csvFileName)
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		errMsg("Failed to parse the provided CSV file.")
	}

	correctCounter := 0
	problems := parseCsvLines(lines)

	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", index+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer) //Does not work with multiple words
		if answer == problem.answer {
			fmt.Println("Correct!")
			correctCounter++
		} else {
			fmt.Printf("Incorrect! The correct answer is: %s\n", problem.answer)
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correctCounter, len(problems))
}
