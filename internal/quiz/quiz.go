package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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
		quizQuestionsAndAnswers.Questions = append(quizQuestionsAndAnswers.Questions, &QuizQuestion)
		log.Printf("New question '%s' & answer '%d' recovered", QuizQuestion.Question, QuizQuestion.Answer)
	}

	return &quizQuestionsAndAnswers, nil
}

func getAnswerToQuestion(question *core.QuizQuestion, buffer *bufio.Reader) (int, error) {
	var err error
	var numberText string
	var value int

	fmt.Printf("What is the answer to %s?\n", question.Question)

	numberText, err = buffer.ReadString('\n')
	if err != nil {
		return 0, err
	}

	value, err = strconv.Atoi(strings.TrimSpace(numberText))
	if err != nil {
		return 0, err
	}

	return value, nil
}

func Results(quizQuestions *core.QuizQuestions) (int, int) {
	var correctAnswers int = 0
	var incorrectAnswers int = 0

	for _, question := range quizQuestions.Questions {
		if question.Answer == question.Response {
			correctAnswers++
		} else {
			incorrectAnswers++
		}
	}

	return correctAnswers, incorrectAnswers
}

func RunQuizGame(quizQuestions *core.QuizQuestions, answer chan core.QuestionAnswer) {
	buffer := bufio.NewReader(os.Stdin)

	for index, question := range quizQuestions.Questions {
		value, err := getAnswerToQuestion(question, buffer)
		if err != nil {
			log.Fatalf("Failed to get answer to question '%s': %v", question.Question, err)
		}
		question.Response = value
		answer <- core.QuestionAnswer{
			QuestionNumber: index,
			Answer:         value,
		}
	}

	// This is a signal to stop the game
	answer <- core.QuestionAnswer{
		QuestionNumber: -1,
		Answer:         -1,
	}
}
