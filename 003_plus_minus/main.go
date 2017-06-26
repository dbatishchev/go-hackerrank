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

	Fraction(data)
}

func Fraction(list []int) {
	zeroes := 0
	positives := 0
	negatives := 0

	for _, v := range list {
		if v > 0 {
			positives++
		} else if v < 0 {
			negatives++
		} else {
			zeroes++
		}
	}

	fmt.Println(fmt.Sprintf("%.6f", float64(positives)/float64(len(list))))
	fmt.Println(fmt.Sprintf("%.6f", float64(negatives)/float64(len(list))))
	fmt.Println(fmt.Sprintf("%.6f", float64(zeroes)/float64(len(list))))
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