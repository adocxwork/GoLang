package main

// file name should end with _test.go
// function name should start with TestXxx.go
// make sure to import testing

// go test -run TestAdd -> single test
// go test -> whole test (includes all files)

// run "go test" from the terminal
// go test main_test.go -> does not works -> main.go is not included

import "testing"


func TestAdd(t *testing.T) {
	result := add(4, 5)
	if result != 9 {
		t.Errorf("add(4, 5) failed, Expected %d, got %d", 9, result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(4, 5)
	if result != -1 {
		t.Errorf("subtract(4, 5) failed, Expected %d, got %d", -1, result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(4, 5)
	if result != 20 {
		t.Errorf("multiply(4, 5) failed, Expected %d, got %d", 20, result)
	}
}

func TestDivide(t *testing.T) {
	result1, err := divide("4", "2")
	if result1 != 2.0 {
		t.Errorf("divide(\"4\", \"2\") failed, Expected %f, got %f", 2.0, result1)
	} else if err != nil {
		t.Errorf("divide(\"4\", \"2\") failed, It returns error like %v", err)
	}
}