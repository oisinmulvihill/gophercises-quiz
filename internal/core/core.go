package core

import "errors"

type QuizQuestion struct {
	Question string
	Answer   int
}

type QuizQuestions struct {
	Questions []QuizQuestion
}

var ErrAnswerNotAnInteger = errors.New("answer is not an integer")
