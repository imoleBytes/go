package utils

import (
	"encoding/csv"
	"os"
)

func ReadCSV(csv_file string) ([][]string, error) {

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
	return records, nil
}
