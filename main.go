package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type GameInfo struct {
	name        string
	description string
	game        func(chan bool)
}

var games map[int]GameInfo = map[int]GameInfo{

	1: {
		"Prime Validator",
		"Is it a prime number?",
		PrimeValidator,
	},
	2: {
		"Baskin Robbins 31",
		"The one who first enters 31 will lose.",
		BaskinRobbins,
	},
}

func main() {

	// print a language selection message
	selectLanguage()

	// print a welcome message
	greetings()

	// let user select a game
	for {

		introMain()

		game, quit := selectGame()

		if quit {
			fmt.Println("Bye.")
			return
		}

		if game != nil {
			quitChan := make(chan bool)
			go game(quitChan)
			<-quitChan
		}

	}

}

func selectGame() (func(chan bool), bool) {

	inputChan := make(chan string)
	doneChan := make(chan bool)

	prompt()
	go readUserInputMain(os.Stdin, inputChan, doneChan)

	select {

	case <-doneChan:
		// fmt.Println("Bye.")
		return nil, true

	case input := <-inputChan:

		k, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			return nil, false
		}

		v, ok := games[k]
		if !ok {
			fmt.Println("Please enter a number in a game list.")
			return nil, false
		}

		fmt.Printf("Starting %s\n", v.name)
		return v.game, false
	}
}

func selectLanguage() {
	fmt.Println("Choose a Language. 언어를 선택하세요.")
}

func greetings() {
	fmt.Println("Welcome to Time Killer!")
	fmt.Println("Remember you can quit whenever you enter Q(q), or ctrl + c.")
}

func introMain() {
	fmt.Println("Please choose a game in a game list below.")
	fmt.Println("------------")
	for k, v := range games {
		fmt.Printf("%d: %s\n", k, v.name)
	}
	fmt.Println("------------")
}

func readUserInputMain(in io.Reader, inputChan chan string, doneChan chan bool) {

	scanner := bufio.NewScanner(in)
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		doneChan <- true
		close(doneChan)
		return
	}

	inputChan <- scanner.Text()
	close(inputChan)

}

func prompt() {
	fmt.Print("-> ")
}
