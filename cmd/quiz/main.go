package main

import (
	"log"
	"os"

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

	quizQuestions, err = quiz.RecoverQuestionsAndAnswers(file)
	if err != nil {
		log.Fatalf("Failed to recover questions and answers: %v", err)
	}

	quiz.RunQuizGame(quizQuestions)
}
