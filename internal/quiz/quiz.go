package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/oisinmulvihill/gophercises-quiz/internal/core"
)

func RecoverQuestionsAndAnswers(source io.Reader, shuffle bool) (*core.QuizQuestions, error) {

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
		// log.Printf("New question '%s' & answer '%d' recovered", QuizQuestion.Question, QuizQuestion.Answer)
	}

	if shuffle {
		fmt.Println("Shuffling questions...")
		rand.Shuffle(len(quizQuestionsAndAnswers.Questions), func(i, j int) {
			quizQuestionsAndAnswers.Questions[i], quizQuestionsAndAnswers.Questions[j] = quizQuestionsAndAnswers.Questions[j], quizQuestionsAndAnswers.Questions[i]
		})
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

	input := strings.TrimSpace(numberText)
	value, err = strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Failed to convert '%s' to a number. The answer is incorrect!\n", input)
		return -1, nil
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
