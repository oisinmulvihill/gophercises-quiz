package settings

import (
	"testing"
)

func TestRecoverFileGiven(t *testing.T) {
	// TestRecover tests the Recover function
	args := []string{"-csv", "test.csv"}

	config := Recover(args)

	if config.QuizQuestionsFile != "test.csv" {
		t.Errorf("expected 'test.csv' but got '%s'", config.QuizQuestionsFile)
	}
}

func TestRecoverDefaultFilename(t *testing.T) {
	args := []string{}

	config := Recover(args)

	if config.QuizQuestionsFile != "problems.csv" {
		t.Errorf("expected 'problems.csv' but got '%s'", config.QuizQuestionsFile)
	}
}
func TestRecoverTimeOut(t *testing.T) {
	args := []string{"-timeout", "10"}

	config := Recover(args)

	if config.TimeOut != 10 {
		t.Errorf("expected '10' TimeOut but got '%v'", config.TimeOut)
	}
}
