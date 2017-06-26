package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	MinMax(data)
}

func MinMax(list []int) {
	min := 0
	max := 0

	sort.Ints(list)

	for _, v := range list[0:4] {
		min += v
	}

	for _, v := range list[1:5] {
		max += v
	}

	fmt.Print(min, max)
}

func readData() ([]int, error) {
	data := make([]int, 5)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}