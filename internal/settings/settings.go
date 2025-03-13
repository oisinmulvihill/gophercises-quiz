package settings

import (
	"flag"
)

type configuration struct {
	QuizQuestionsFile string
	TimeOut           int
	Shuffle           bool
}

func Recover(arguments []string) *configuration {

	flags := flag.FlagSet{}
	quizQuestionsFile := flags.String("csv", "problems.csv", "A csv formated file in the format of 'question,answer'")
	timeOut := flags.Int("timeout", 30, "The time in seconds to complete the quiz")
	shuffle := flags.Bool("shuffle", false, "Use to randomse the questions.")
	flags.Parse(arguments)

	config := configuration{
		QuizQuestionsFile: *quizQuestionsFile,
		TimeOut:           *timeOut,
		Shuffle:           *shuffle,
	}

	return &config
}
