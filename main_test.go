package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_IsPrime(t *testing.T) {

	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"no prime", 1, false, "1 is not prime, by definition!"},
		{"no prime", 8, false, "8 is not a prime number, because it is divisible by 2!"},
		{"no prime", -3, false, "Negative numbers are not prime, by definition!"},
		{"no prime", 0, false, "0 is not prime, by definition!"},
	}

	for _, e := range primeTests {

		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true, but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false, but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s expected %s, but got %s", e.name, e.msg, msg)
		}
	}

}

func Test_Prompt(t *testing.T) {

	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() from our read pipe
	out, _ := io.ReadAll(r)

	// perform test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}

}

func Test_Intro(t *testing.T) {

	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close writer
	_ = w.Close()

	// reset os.Stdout to wht it was before
	os.Stdout = oldOut

	// read the output of our intro() from our read pipe
	out, _ := io.ReadAll(r)

	// perform test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("Incorrect intro text: got %s", string(out))
	}

}

func Test_CheckNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "quit", input: "Q", expected: ""},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "prime", input: "2", expected: "2 is a prime number!"},
		{name: "no prime", input: "4", expected: "4 is not a prime number, because it is divisible by 2!"},
		{name: "no prime", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected: %s, but got: %s", e.name, e.expected, res)
		}
	}

}
