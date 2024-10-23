package utils

import (
	"encoding/csv"
	"errors"
	"os"
)

func ReadCSV(csv_file string) ([][]string, error) {

	file, err := os.Open(csv_file)

	if err != nil {
		return nil, errors.New("error openinf file o")
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
