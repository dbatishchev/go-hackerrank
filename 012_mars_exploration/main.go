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

	ExploreMars(input)
}

func ExploreMars(input []byte) {
	sos := "SOS"
	counter := 0
	for i, v := range input {
		correctChar := sos[i % 3]
		if v != correctChar {
			counter++
		}
	}

	fmt.Println(counter);
}

func readData() ([]byte, error) {
	var input []byte

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return nil, err
	}

	return input, nil
}