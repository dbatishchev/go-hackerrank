package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	size, input, k, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	result := Encrypt(size, input, k)

	fmt.Println(string(result))
}

func Encrypt(size int, input []byte, k int) []byte {
	var result = make([]byte, len(input))
	var inputAsString = string(input)
	const lettersCountInEnglish = byte(26)

	for i, v := range inputAsString {
		byteValue := byte(v)
		if (unicode.IsLetter(v)) {
			var base byte
			if (unicode.IsLower(v)) {
				base = byte('a')
			} else {
				base = byte('A')
			}
			encryptedChar := (byteValue + byte(k) - base) % lettersCountInEnglish + base
			result[i] = encryptedChar
		} else {
			result[i] = byteValue
		}
	}

	return result
}

func readData() (size int, input []byte, k int, error error) {
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%s", &input)
	if err != nil {
		return
	}

	_, err = fmt.Scanf("%d", &k)
	if err != nil {
		return
	}

	return
}