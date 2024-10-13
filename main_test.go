package main

import (
	"testing"
)

type value struct {
	input    [][]string
	expected int
}

var values = []value{
	{input: [][]string{{"5+5", "10"}}, expected: 1},
}

func TestQuiz(t *testing.T) {

	got, err := quizTest(values[0].input)
	want := 1
	var wantErr error

	if got != want || wantErr != err {
		t.Errorf("quizTest() = %d, %v; want %d, %v", got, err, want, wantErr)
	}

}
