package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/oisinmulvihill/gophercises-quiz/internal/core"
)

func RecoverQuestionsAndAnswers(source io.Reader) (*core.QuizQuestions, error) {

	quizQuestionsAndAnswers := core.QuizQuestions{}

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
			errorMessage := fmt.Errorf("cannot convert answer to integer '%v' for question '%s' because %w", record[1], record[0], err)
			log.Println(errorMessage)
			return &quizQuestionsAndAnswers, fmt.Errorf("%s: %w", errorMessage, core.ErrAnswerNotAnInteger)
		}

		QuizQuestion := core.QuizQuestion{
			Question: record[0],
			Answer:   answer,
		}
		quizQuestionsAndAnswers.Questions = append(quizQuestionsAndAnswers.Questions, QuizQuestion)
		log.Printf("New question '%s' & answer '%d' recovered", QuizQuestion.Question, QuizQuestion.Answer)
	}

	return &quizQuestionsAndAnswers, nil
}
