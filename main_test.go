package main

import (
	"io"
	"os"
	"testing"
)

func Test_isPrime(t *testing.T) {

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
