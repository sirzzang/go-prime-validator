package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var primeValidatorDescription = `Enter a whole number.
We'll tell you if it is a prime number or not.
Enter Q(q) or ctrl + c to quit.`

func PrimeValidator(quitChan chan bool) {

	introPrime()

	doneChan := make(chan bool)
	go readUserInputPrime(os.Stdin, doneChan)
	<-doneChan
	close(doneChan)

	// say goodbye
	fmt.Println("Going back to main.")
	quitChan <- true
	close(quitChan)

}

func readUserInputPrime(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbersPrime(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbersPrime(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func introPrime() {
	fmt.Println(primeValidatorDescription)
	prompt()
}

func isPrime(n int) (bool, string) {
	// 0 and 1 is not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number, because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)

}
