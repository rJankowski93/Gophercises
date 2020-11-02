package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	csvFileName := flag.String("csvFileName", "data.csv", "Name of csv file")
	flag.Parse()
	csvfile, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		os.Exit(1)
	}
	reader := csv.NewReader(csvfile)

	numberCorrectAnswers := 0
	numberIncorrectAnswers := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var answer string
		fmt.Printf("%s = ", record[0])
		fmt.Scanln(&answer)
		if answer == record[1] {
			fmt.Println("Correct")
			numberCorrectAnswers++
		} else {
			fmt.Printf("Wrong, correct answer: %s\n", record[1])
			numberIncorrectAnswers++
		}
	}
	fmt.Printf("Number of correct answers: %v\n", numberCorrectAnswers)
	fmt.Printf("Number of incorrect answers: %v\n", numberIncorrectAnswers)
}
