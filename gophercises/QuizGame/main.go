package main

import (
	"flag"
	"fmt"
	"os"
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

	score := 0
	done := make(chan bool)
	answerch := make(chan bool)
	var user_answer string

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
