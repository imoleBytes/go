package utils

import (
	"fmt"
)

func Ask(num int, question, answer string) int {
	var user_answer string

	fmt.Printf("Problem #%v: %v = ", num, question)
	fmt.Scanf("%s\n", &user_answer)

	if user_answer == answer {
		return 1
	}
	return 0
}
