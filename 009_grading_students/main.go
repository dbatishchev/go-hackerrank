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

	RoundGrades(data)
}

func RoundGrades(list []int) {
	roundedGrades := make([]int, len(list))
	for i, g := range list {
		println(math.Mod(float64(g), 5.0))
		if g < 38 {
			roundedGrades[i] = g
		} else if math.Mod(float64(g), 5.0) >= 3 {
			roundedGrades[i] = g + (5 - int(math.Mod(float64(g), 5.0)))
		} else {
			roundedGrades[i] = g
		}
	}

	for _, v := range roundedGrades {
		fmt.Println(v)
	}
}

func readData() ([]int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int, length)

	for i := 0; i < length; i++ {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}