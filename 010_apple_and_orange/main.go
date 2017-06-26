package main

import (
	"fmt"
	"log"
)

func main() {
	s, t, a, b, apples, oranges, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	successedApples, successedOranges := CountFruits(s, t, a, b, apples, oranges)

	fmt.Println(successedApples)
	fmt.Println(successedOranges)
}

func CountFruits(s, t, a, b int, apples, oranges []int) (int, int) {
	return countSuccesses(a, s, t, apples), countSuccesses(b, s, t, oranges)
}

func countSuccesses(startPoint, s, t int, fruits []int) int {
	successes := 0
	for _, distance := range fruits {
		pointOfFall := startPoint + distance
		if pointOfFall >= s && pointOfFall <= t {
			successes++
		}
	}

	return successes
}

func readData() (s, t, a, b int, apples, oranges []int, err error) {
	_, err = fmt.Scanf("%d", &s)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%d", &t)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%d", &a)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		return
	}

	var m int
	_, err = fmt.Scanf("%d", &m)
	if err != nil {
		return
	}

	var n int
	_, err = fmt.Scanf("%d", &n)
	if err != nil {
		return
	}

	apples = make([]int, m)
	for i := 0; i < m; i++ {
		_, err = fmt.Scanf("%d", &apples[i])
		if err != nil {
			return
		}
	}

	oranges = make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scanf("%d", &oranges[i])
		if err != nil {
			return
		}
	}

	return
}