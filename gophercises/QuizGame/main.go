package main

import (
	"flag"
	"fmt"
	"gophercises/quizGame/utils"
)

func main() {

	csvfile := flag.String("file", "problems.csv", "path to csv file")

	flag.Parse()

	records, err := utils.ReadCSV(*csvfile)
	if err != nil {
		fmt.Println("eroor o!")
	}

	total := len(records)

	score := 0
	for index := range records {
		record := records[index]

		que := record[0]
		ans := record[1]

		score += utils.Ask(index+1, que, ans)
	}
	fmt.Printf("You scored %v out of %v.\n", score, total)
}
