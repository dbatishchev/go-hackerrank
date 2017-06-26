package main

import (
	"fmt"
	"log"
)

func main() {
	size, input, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	result := CountSteps(size, input)

	fmt.Println(result)
}

func CountSteps(size int, input []byte) int {
	tripletIndexes := []int{}

	for i := 0; i < len(input); {
		if i + 3 > len(input) {
			break;
		}
		if string(input[i : i + 3]) == "010" {
			tripletIndexes = append(tripletIndexes, i)
			i = i + 3
		} else {
			i = i + 1
		}
	}

	return len(tripletIndexes)
}

func readData() (size int, input []byte, err error) {
	_, err = fmt.Scanf("%d", &size)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%s", &input)
	if err != nil {
		return
	}

	return
}