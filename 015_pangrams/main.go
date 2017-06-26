package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"unicode"
)

func main() {
	input, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	IsPangram(input)
}

func IsPangram(input string) {
	m := make(map[rune]int)

	for _, elem := range input {
		if (unicode.IsLetter(elem)) {
			m[unicode.ToLower(elem)] = 1
		}
	}

	if (len(m) == 26) {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}

func readData() (string, error) {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
	}

	return input, nil
}