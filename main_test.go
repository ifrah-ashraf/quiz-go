package main

import (
	"context"
	"strings"
	"testing"
	"time"
)

var testRecords = [][]string{
	{"5+5", "10"},
	{"11+2", "13"},
	{"13+2", "15"},
	{"6*7", "42"},
}

func TestQuizTest(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)

	inputData := "10\n13\n15\n42"

	defer cancelFunc()

	val := strings.NewReader(inputData)

	got, err := quizTest(ctx, testRecords, val)

	want := 4

	if got != want || err != nil {
		t.Errorf("Expected score %d got score %d", want, got)
	}

}

func TestInvalidValue(t *testing.T) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)

	inputData := "xyz\npqr\nlp\npop"

	defer cancelFunc()

	val := strings.NewReader(inputData)

	got, err := quizTest(ctx, testRecords, val)

	want := 0

	if got != want || err != nil {
		t.Errorf("Expected score %d got score %d", want, got)
	}
}
