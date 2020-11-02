package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csvFileName", "data.csv", "Name of csv file")
	timeForQuestion := flag.Int64("timeForQuestion", 10, "Time for question")
	flag.Parse()

	lines := getLinesFromFile(csvFileName)
	problems := parseLines(lines)
	shuffleSlice(problems)
	numberCorrectAnswers := 0

	fmt.Println("Press enter if do you ready to start quiz")
	fmt.Scanln()

	answerCh := make(chan string, 1)
	for _, problem := range problems {
		var answer string
		fmt.Printf("%s = ", problem.question)

		go func() {
			fmt.Scanln(&answer)
			answerCh <- answer
		}()
		select {
		case answer := <-answerCh:

			if answer == problem.answer {
				fmt.Println("Correct")
				numberCorrectAnswers++
			} else {
				fmt.Printf("Wrong, correct answer: %s\n", answer, problem.answer)
			}
		case <-time.After(time.Duration(*timeForQuestion) * time.Second):
			fmt.Println("Wrong. Out of time")
		}
	}

	fmt.Printf("Number of correct answers: %v\n", numberCorrectAnswers)
	fmt.Printf("Number of total answers: %v\n", len(problems))

}

func shuffleSlice(problems []problem) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
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
