package utcase

import (
	"testing"
)

func Test_Add(t *testing.T) {
	expected := 5

	if actual := Add(3, 2); actual != expected {
		t.Errorf("Add() = %v, want %v", actual, expected)
	}
}

func Test_Sum(t *testing.T) {
	expect := 6

	if result, _ := Sum(1, 2, 3); result != expect {
		t.Errorf("Sum() failed, expected: %v, got: %v", expect, result)
	} else {
		t.Log("Sum() passed.")
	}
}
