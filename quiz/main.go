package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "Name of the csv file having problems")
	timeFlag := flag.Int("duration", 10, "time limit to finish the game")
	flag.Parse()
	csvFileName := *csvFile // storing the file name instead of address
	timeLimit := *timeFlag

	// opening the file
	file, err := os.Open(csvFileName)
	if err != nil {
		fmt.Println("error in opening file")
	}
	// Read the quiz
	quiz := csvFileReader(file)

	// parse questions
	problems := questionParser(quiz)
	result := runGame(problems, timeLimit)
	fmt.Printf("you answered %d correct answers out of %d questions ", result, len(problems))
	// close all the go routines and channels
}

func runGame(problems []Problem, timeLimit int) int {
	answerChannel := make(chan string)
	correctAnswerCounter := 0
	// wait for the player to press enter, before starting the game.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to run the game")
	text, err := reader.ReadString('\n')

	if err != nil {
		panic("error reading the input")
	}

	text = strings.TrimSpace(text)
	if len(text) == 0 {
		// When user presses enter, start the game
		// setup the timer.
		timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	gameLoop:
		for index, problem := range problems {

			question := problem.question
			answer := problem.answer

			fmt.Printf("Problem %d %s \n", index+1, question)
			go inputAnswers(answerChannel)
			select {
			case userInput := <-answerChannel:
				isTrue := answerChecker(userInput, answer)
				if isTrue {
					correctAnswerCounter++
				}
			case <-timer.C:
				fmt.Println("Time Up !")
				break gameLoop
			}
		}
	}
	fmt.Println("You pressed the wrong key")
	return correctAnswerCounter
}

func answerChecker(userInput, answer string) bool {
	if userInput == answer {
		return true
	}
	return false
}

func inputAnswers(answerChannel chan string) {
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("error in taking input of value")
	}
	answerChannel <- userInput
}

func csvFileReader(fileName *os.File) [][]string {
	csvReader := csv.NewReader(fileName)
	quiz, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}
	return quiz
}

func questionParser(quiz [][]string) []Problem {
	problems := make([]Problem, len(quiz))
	for index, entry := range quiz {
		problems[index] = Problem{
			question: entry[0],
			answer:   entry[1],
		}
	}
	return problems
}
