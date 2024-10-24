package main

import (
	"flag"
	"fmt"
	"gophercises/quizGame/utils"
	"os"
	"strings"
)

func main() {

	csvfile := flag.String("file", "problems.csv", "path to csv file, USAGE: -file=test.csv")

	flag.Parse()

	records, err := readCSV(*csvfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	total := len(records)

	score := 0
	for index := range records {
		record := records[index]

		que := record[0]
		ans := strings.TrimSpace(record[1])

		score += utils.Ask(index+1, que, ans)
	}
	fmt.Printf("You scored %v out of %v.\n", score, total)
}
