package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	aliceGrades := [3]int{}
	bobGrades := [3]int{}

	aliceScore := 0
	bobScore := 0

	io := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &aliceGrades[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &bobGrades[i])
	}

	for i := 0; i < 3; i++ {
		if aliceGrades[i] > bobGrades[i] {
			aliceScore++
		} else if aliceGrades[i] < bobGrades[i] {
			bobScore++
		}
	}

	fmt.Print(aliceScore, bobScore)
}
