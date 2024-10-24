package main

import (
	"flag"
	"fmt"
	"gophercises/quizGame/utils"
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

	donech := make(chan bool)
	score := 0
	go func(done chan bool) {

	outer:
		for index := range records {
			select {
			case <-done:
				break outer
			default:
				record := records[index]

				que := record[0]
				ans := strings.TrimSpace(record[1])

				score += utils.Ask(index+1, que, ans)
			}

		}
	}(donech)
	sleepfor(*limit, donech)

	fmt.Printf("You scored %v out of %v.\n", score, total)
}

func sleepfor(t int, done chan bool) {
	time.Sleep(time.Duration(t) * time.Second)
	done <- true
}
