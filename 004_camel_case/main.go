package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	input, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	result := CamelCaseCount(input)

	fmt.Print(result)
}

func CamelCaseCount(input string) int {
	if len(input) == 0 {
		return 0
	}

	wordsCount := 1

	for _, elem := range input {
		if unicode.IsUpper(elem) {
			wordsCount++
		}
	}

	return wordsCount
}

func readData() (string, error) {
	var input string

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return "", err
	}

	return input, nil
}