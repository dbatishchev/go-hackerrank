package main

import (
	"fmt"
	"log"
)

func main() {
	input, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	result := Reduce(input)

	if len(result) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Println(string(result))
	}
}

func Reduce(input []byte) []byte {
	if len(input) < 2 {
		return input
	}


	for i := 1; i < len(input); i++ {
		if input[i] == input[i - 1] {
			input = append(input[0:i-1], input[i+1:]...)
			i = 0
		}
	}

	return input
}

func readData() ([]byte, error) {
	var input []byte

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return nil, err
	}

	return input, nil
}