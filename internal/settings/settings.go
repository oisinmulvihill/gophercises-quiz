package settings

import (
	"flag"
)

type configuration struct {
	QuizQuestionsFile string
	TimeOut           int
}

func Recover(arguments []string) *configuration {

	flags := flag.FlagSet{}
	quizQuestionsFile := flags.String("csv", "problems.csv", "A csv formated file in the format of 'question,answer'")
	flags.Parse(arguments)

	config := configuration{
		QuizQuestionsFile: *quizQuestionsFile,
	}

	return &config
}
