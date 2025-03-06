package quiz

import (
	"strings"
	"testing"
)

func TestRecoverQuizQuestions(t *testing.T) {

	source := strings.NewReader(`5+5,10
7+7,14`)

	quizQuestions, err := RecoverQuizQuestions(source)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(quizQuestions.Questions) != 2 {
		t.Errorf("expected 4 questions but got %d", len(quizQuestions.Questions))
	}

	if quizQuestions.Questions[0].Question != "5+5" {
		t.Errorf("expected '5+5' but got '%s'", quizQuestions.Questions[0].Question)
	}
	if quizQuestions.Questions[0].Answer != 10 {
		t.Errorf("expected '10' but got '%d'", quizQuestions.Questions[0].Answer)
	}

	if quizQuestions.Questions[1].Question != "7+7" {
		t.Errorf("expected '7+7' but got '%s'", quizQuestions.Questions[0].Question)
	}
	if quizQuestions.Questions[1].Answer != 14 {
		t.Errorf("expected '14' but got '%d'", quizQuestions.Questions[0].Answer)
	}
}

func TestRecoverQuizQuestionsNoRows(t *testing.T) {

	source := strings.NewReader("")

	quizQuestions, err := RecoverQuizQuestions(source)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(quizQuestions.Questions) != 0 {
		t.Errorf("expected 0 questions but got %d", len(quizQuestions.Questions))
	}
}
func TestRecoverQuizQuestionsAnswerFailedToConvertToInt(t *testing.T) {

	source := strings.NewReader("2+2,a")

	_, err := RecoverQuizQuestions(source)

	expectedError := "cannot convert answer to integer: 'a' for question: 2+2"

	if err.Error() != expectedError {
		t.Errorf("unexpected error: %v", err)
	}
}
