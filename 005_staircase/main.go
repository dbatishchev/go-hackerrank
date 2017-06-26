package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	Staircase(data)
}

func Staircase(size int) {
	for i := 1; i <= size; i++ {
		fmt.Println(strings.Repeat(" ", size - i) + strings.Repeat("#", i))
	}
}

func readData() (int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return 0, err
	}

	return length, nil
}