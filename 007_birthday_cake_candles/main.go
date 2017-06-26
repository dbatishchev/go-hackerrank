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

	fmt.Println(ContOfBlowedCandles(data))
}

func ContOfBlowedCandles(list []int) int {
	max := 0
	count := 0

	for _, v := range list {
		if v > max {
			max = v
			count = 1
		} else if v == max {
			count++
		}
	}

	return count
}

func readData() ([]int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int, length)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}