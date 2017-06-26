package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	firstInput, secondInput, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	result := AnagramCount(firstInput, secondInput)

	fmt.Println(result)
}

func AnagramCount(firstInput, secondInput string) int {
	const lettersInEnglishAlphabet = 26
	var firstLetterPosition = 'a'
	var firstStringLettersCount [lettersInEnglishAlphabet]int
	var secondStringLettersCount [lettersInEnglishAlphabet]int

	for _, c := range firstInput {
		firstStringLettersCount[c % firstLetterPosition]++ // rune is an alias for int32
	}

	for _, c := range secondInput {
		secondStringLettersCount[c % firstLetterPosition]++
	}

	result := 0

	for i := 0; i < lettersInEnglishAlphabet; i++ {
		result += int(math.Abs(float64(firstStringLettersCount[i] - secondStringLettersCount[i])))
	}

	return result
}

func readData() (firstInput, secondInput string, err error) {
	_, err = fmt.Scanf("%s", &firstInput)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%s", &secondInput)
	if err != nil {
		return
	}

	return
}