package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", Difference(data))
}

func Difference(matrix [][]int) int {
	length := len(matrix)
	primaryDiagonal := 0
	secondaryDiagonal := 0

	for i := 0; i < length; i++ {
		primaryDiagonal += matrix[i][i]
		secondaryDiagonal += matrix[i][length - i - 1]
	}

	return int(math.Abs(float64(primaryDiagonal - secondaryDiagonal)))
}

func readData() ([][]int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([][]int, length)

	for i := 0; i < length; i++ {
		row := make([]int, length)
		for j := 0; j < length; j++ {
			_, err := fmt.Scanf("%d", &row[j])
			if err != nil {
				return nil, err
			}
		}
		data[i] = row
	}

	return data, nil
}