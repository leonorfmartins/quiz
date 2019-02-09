package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var csvPath = flag.String("csvPath", "./problems.csv", "Path to CSV file")

func main() {
	flag.Parse()
	// exception! We can deference this because there is a need of the flags package
	// to allocate the string for us somewhere
	result := readCSVFile(*csvPath)
	// rightAnswers accumulates how many right questions the user gotvar RightAnswers int
	var numberRightAnswers int
	for _, qna := range result {
		rightGuess := qna.displayAndCountAnswers(os.Stdin)
		if rightGuess == true {
			numberRightAnswers++
		}
	}
	displayResultMessage(numberRightAnswers, len(result))
}

func displayResultMessage(numberRightAnswers int, numberQuestions int) {
	if numberRightAnswers == numberQuestions {
		fmt.Printf("You answered all the questions correctly! Congrats!")
	} else {
		fmt.Printf("You got %v right answers out of %v \n", numberRightAnswers, numberQuestions)
	}
}

// QuestionsAndAnswers has questions and answers mapped
type QuestionsAndAnswers struct {
	Question string
	Answer   int
}

func readCSVFile(fileName string) []QuestionsAndAnswers {
	fileReader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(fileReader)
	var entries []QuestionsAndAnswers
	for {
		entry, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		answer, err := strconv.Atoi(entry[1])
		if err != nil {
			log.Fatal(err)
		}
		qna := QuestionsAndAnswers{
			Question: entry[0],
			Answer:   answer,
		}
		entries = append(entries, qna)
	}
	fileReader.Close()
	return entries
}

func (qna QuestionsAndAnswers) displayAndCountAnswers(r io.Reader) bool {
	var guess int
	for {
		fmt.Printf("%v\n", qna.Question)
		var userInput string
		_, err := fmt.Fscan(r, &userInput)
		if err != nil {
			log.Fatal(err)
		}
		guess, err = strconv.Atoi(userInput)
		if err != nil {
			log.Printf("Please input a number as an answer")
			continue
		}
		break
	}
	if guess == qna.Answer {
		fmt.Printf("right on, right answer!\n")
		return true
	}
	fmt.Printf("oh no, wrong answer :( \n")
	return false
}
