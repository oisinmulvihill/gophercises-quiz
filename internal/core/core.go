package core

import "errors"

type QuizQuestion struct {
	Question string
	Answer   int
	Response int
}

type QuestionAnswer struct {
	QuestionNumber int
	Answer         int
}

type QuizQuestions struct {
	Questions []*QuizQuestion
}

var ErrAnswerNotAnInteger = errors.New("answer is not an integer")
