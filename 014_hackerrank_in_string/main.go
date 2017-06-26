package main

import (
	"fmt"
	"log"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	ContainsHackerRank(data)
}

func ContainsHackerRank(data []string) {
	const hackerrank = "hackerrank"
	for _, str := range data {
		isValid := false
		currentChar := 0
		for _, charInInputString := range str {
			if charInInputString == rune(hackerrank[currentChar]) {
				currentChar = currentChar + 1
				if currentChar == len(hackerrank) {
					isValid = true
					break
				}
			}
		}
		if isValid {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func readData() ([]string, error) {
	var size int
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		return nil, err
	}

	data := make([]string, size)

	for i := 0; i < size; i++ {
		var inputString string
		_, err := fmt.Scanf("%s", &inputString)
		if err != nil {
			return nil, err
		}
		data[i] = inputString
	}

	return data, nil
}