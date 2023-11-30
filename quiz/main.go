package main

// TODO: Use logger for logging.
import (
	"encoding/csv"
	"fmt"
	"os"
)

type Quiz struct {
	questions []*Questions
	fileName  string
}

type Questions struct {
	question      string
	correctAnswer string
	userAnswer    string
}

func (q *Quiz) StartQuiz() error {
	fmt.Println("QUIZ STARTED")
	err := q.createQuiz()
	if err != nil {
		return err
	}
	for index := 0; index < len(q.questions); index++ {
		fmt.Println("Question ", q.questions[index].question)
		fmt.Print("Your Answer ")
		fmt.Scan(&q.questions[index].userAnswer)
	}
	q.endQuiz()
	return nil
}

func (q *Quiz) endQuiz() {
	correct := 0
	total := len(q.questions)
	for index := 0; index < len(q.questions); index++ {
		correctAnswer := q.questions[index].correctAnswer
		userAnswer := q.questions[index].userAnswer
		if correctAnswer == userAnswer {
			correct += 1
		}
	}
	fmt.Printf("You got %d correct out of %d questions", correct, total)
	return
}

func (q *Quiz) createQuiz() error {
	records, err := q.readQuiz()
	if err != nil {
		return err
	}

	for index := 0; index < len(records); index++ {
		question := &Questions{
			records[index][0],
			records[index][1],
			"",
		}

		q.questions = append(q.questions, question)
	}

	return nil
}

func (q *Quiz) readQuiz() ([][]string, error) {
	q.fileName = "problems.csv"
	if len(os.Args) > 2 {
		q.fileName = os.Args[2]
	}
	fmt.Println("Quiz File", q.fileName)

	fd, err := os.Open(q.fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return nil, err
	}

	csvReader := csv.NewReader(fd)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading from file")
		return nil, err
	}
	return records, nil
}

func main() {
	quiz := &Quiz{}
	err := quiz.StartQuiz()
	if err != nil {
		fmt.Println(err)
	}
}
