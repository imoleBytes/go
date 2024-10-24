package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvfile := flag.String("file", "problems.csv", "path to csv file, USAGE: -file=test.csv")
	limit := flag.Int("limit", 10, "The time limit for the quiz in seconds, USAGE: -limit=5")

	flag.Parse()

	records, err := readCSV(*csvfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	total := len(records)

	var score int = 0
	done := make(chan bool)
	answerch := make(chan bool)
	var user_answer string

	var play string

	for {
		fmt.Println("Questions Ready! Are you ready to play? (Y/N) ")
		fmt.Scanf("%s\n", &play)

		play = strings.ToLower(play)
		if play == "n" {
			fmt.Println("You missing out!, Bye")
			os.Exit(0)
		} else if play != "y" {
			fmt.Println("Invalid response!: Enter either y or n")
		} else {
			fmt.Printf("You have %v seconds, make it count! ENTER to START", *limit)
			fmt.Scanln()
			break
		}

	}

	go sleepfor(*limit, done)

outerLoop:
	for index, problem := range records {
		fmt.Printf("Problem #%v: %v = ", index+1, problem.que)

		go func() {
			fmt.Scanf("%s\n", &user_answer)
			answerch <- true
			if user_answer == problem.ans {
				score++
			}

		}()
		select {
		case <-done:
			fmt.Println("\ntime up!!!!")
			break outerLoop
		case <-answerch:
			continue
		}
	}
	fmt.Printf("You scored %v out of %v.\n", score, total)

}

func sleepfor(t int, done chan bool) {
	time.Sleep(time.Duration(t) * time.Second)
	done <- true
}
