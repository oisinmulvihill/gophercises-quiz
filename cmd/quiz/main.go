package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oisinmulvihill/gophercises-quiz/internal/core"
	"github.com/oisinmulvihill/gophercises-quiz/internal/quiz"
	"github.com/oisinmulvihill/gophercises-quiz/internal/settings"
)

func init() {
	log.Println("Initializing Quiz...")
}

func main() {

	config := settings.Recover(os.Args[1:])

	file, err := os.Open(config.QuizQuestionsFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("The file '%s' does not exist", config.QuizQuestionsFile)
		} else {
			log.Fatal(err)
		}
	}

	var quizQuestions *core.QuizQuestions
	answerChannel := make(chan core.QuestionAnswer)

	quizQuestions, err = quiz.RecoverQuestionsAndAnswers(file)
	if err != nil {
		log.Fatalf("Failed to recover questions and answers: %v", err)
	}

	go quiz.RunQuizGame(quizQuestions, answerChannel)

game:
	for {
		select {
		case response := <-answerChannel:
			if response.QuestionNumber == -1 {
				fmt.Println("Quiz complete!")
				break game
			} else {
				q := quizQuestions.Questions[response.QuestionNumber]
				q.Response = response.Answer
			}

		case <-time.After(10 * time.Second):
			log.Println("Time's up!")
			break game
		}
	}

	correctAnswers, incorrectAnswers := quiz.Results(quizQuestions)
	fmt.Printf("You got %d correct answers and %d incorrect answers\n", correctAnswers, incorrectAnswers)
}
