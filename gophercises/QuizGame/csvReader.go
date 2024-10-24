package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func readCSV(csv_file string) ([]problem, error) {

	file, err := os.Open(csv_file)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	total := len(records)

	problems := make([]problem, total)

	for index, record := range records {
		problems[index] = problem{
			que: strings.TrimSpace(record[0]),
			ans: strings.TrimSpace(record[1]),
		}
	}

	return problems, nil
}

type problem struct {
	que string
	ans string
}
