package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFileName := flag.String("csvFileName", "data.csv", "Name of csv file")
	flag.Parse()
	lines := getLinesFromFile(csvFileName)
	problems := parseLines(lines)
	numberCorrectAnswers := 0

	for _, problem := range problems {
		var answer string
		fmt.Printf("%s = ", problem.question)
		fmt.Scanln(&answer)
		if answer == problem.answer {
			fmt.Println("Correct")
			numberCorrectAnswers++
		} else {
			fmt.Printf("Wrong, correct answer: %s\n", problem.answer)
		}
	}

	fmt.Printf("Number of correct answers: %v\n", numberCorrectAnswers)
	fmt.Printf("Number of total answers: %v\n", len(problems))
}

func getLinesFromFile(csvFileName *string) [][]string {
	csvfile, err := os.Open(*csvFileName)
	handleError(err, "Couldn't open the csv file")
	reader := csv.NewReader(csvfile)
	lines, err := reader.ReadAll()
	handleError(err, "Couldn't open the csv file")
	return lines
}

func handleError(err error, errMsg string) {
	if err != nil {
		log.Fatalln(errMsg, err)
		os.Exit(1)
	}
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] =
			problem{
				line[0], line[1],
			}
	}
	return problems
}

type problem struct {
	question string
	answer   string
}
