package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

type QuizQuestion struct {
	Question string
	Answer   int
}

type QuizQuestions struct {
	Questions []QuizQuestion
}

func RecoverQuizQuestions(source io.Reader) (*QuizQuestions, error) {

	quizQuestionsAndAnswers := QuizQuestions{}

	reader := csv.NewReader(source)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return &quizQuestionsAndAnswers, err
		}

		answer, err := strconv.Atoi(record[1])
		if err != nil {
			err = fmt.Errorf("cannot convert answer to integer: '%v' for question: %s", record[1], record[0])
			log.Fatal(err)
			return &quizQuestionsAndAnswers, err
		}

		QuizQuestion := QuizQuestion{
			Question: record[0],
			Answer:   answer,
		}
		quizQuestionsAndAnswers.Questions = append(quizQuestionsAndAnswers.Questions, QuizQuestion)
		log.Printf("New question '%s' & answer '%d' recovered", QuizQuestion.Question, QuizQuestion.Answer)
	}

	return &quizQuestionsAndAnswers, nil
}
