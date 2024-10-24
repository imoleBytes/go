package main

import (
	"flag"
	"fmt"
	"gophercises/quizGame/utils"
	"os"
)

func main() {

	csvfile := flag.String("file", "problems.csv", "path to csv file, USAGE: -file=test.csv")
	// limit := flag.Int("limit", 10, "The time limit for the quiz in seconds, USAGE: -limit=5")

	flag.Parse()

	records, err := readCSV(*csvfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	total := len(records)

	score := 0

	for index, problem := range records {
		score += utils.Ask(index+1, problem.que, problem.ans)
	}

	fmt.Printf("You scored %v out of %v.\n", score, total)
}

// func sleepfor(t int, done chan bool) {
// 	time.Sleep(time.Duration(t) * time.Second)
// 	done <- true
// }
