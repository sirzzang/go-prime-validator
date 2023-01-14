package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var cnt int = 0
var lastNum int = 0

func main() {

	intro()

	doneChan := make(chan bool)
	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	fmt.Println("Goodbye.")
	fmt.Println("Total cnt:", cnt)

}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		msg, done := checkNumbers(scanner)
		fmt.Println(msg)

		if done {
			doneChan <- true
			return
		}

		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {

	scanner.Scan()
	line := scanner.Text()

	// quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	nums, err := stringToNumbers(line)
	if err != nil {
		return err.Error(), false
	}

	// TODO: custom
	if len(nums) < 1 || len(nums) > 3 {
		fmt.Println("Please enter 1 or more but less than 3 numbers.")
		return "", false
	}

	// TODO: check if numbers are in ascending order and successive
	for i, num := range nums {
		if num != lastNum+i+1 {
			fmt.Printf("Please make sure you enter %d in a right place.\n", lastNum+i+1)
		return "", false
		}
	}

	msg, done := doGame(nums)

	return msg, done

}

func doGame(nums []int) (string, bool) {

	cnt++

	for _, num := range nums {
		// lose when 31 in nums
		if num == 31 {
			return "You lose!", true
		}
	}

	lastNum = nums[len(nums)-1]

	return fmt.Sprintf("%v. Continue game.", nums), false
}

func stringToNumbers(line string) ([]int, error) {

	// TODO: custom
	nums := make([]int, 0, 3)

	for _, v := range strings.Fields(line) {
		i, err := strconv.Atoi(v)
		if err != nil || i < 1 {
			return nil, fmt.Errorf("Please enter positive integers.")
		}

		nums = append(nums, i)
	}
	return nums, nil
}

func intro() {
	fmt.Println("Baskin Robbins 31")
	fmt.Println("------------")
	fmt.Println("Enter one or more than one numbers.")
	fmt.Println("Leave a space between numbers when entering two or more numbers.")
	fmt.Println("All players' input numbers should be successive.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}
