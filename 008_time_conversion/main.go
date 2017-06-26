package main

import (
	"fmt"
	"log"
)

func main() {
	hour, minutes, seconds, period, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	result := ConvertToMilitaryTime(hour, minutes, seconds, period)

	fmt.Print(result)
}

func ConvertToMilitaryTime(hour, minutes, seconds int, period string) string {
	if period == "PM" && hour != 12 {
		hour += 12
	}

	if period == "AM" && hour == 12 {
		hour = 0
	}

	return fmt.Sprintf("%02d:%02d:%02d", hour, minutes, seconds)
}

func readData() (int, int, int, string, error) {
	var hour int
	var minutes int
	var seconds int
	var period string

	_, err := fmt.Scanf("%d:%d:%d%s", &hour, &minutes, &seconds, &period)
	if err != nil {
		return 0, 0, 0,"", err
	}

	return hour, minutes, seconds, period, nil
}