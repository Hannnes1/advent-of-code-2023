package main

import (
	"fmt"
	"unicode"
)

func main() {
	result := part1(input)

	fmt.Printf("Result of part 1: %d\n", result)
}

func part1(input string) int {
  // Add a new line to the end, so that the last line is included in the sum as well.
  input += "\n"

	sum := 0

	firstDigit := 0
	lastDigit := 0

	foundDigit := false

	for _, char := range input {
		if unicode.IsDigit(char) {
			// Convert ASCII to int by subtracting decimal value of '0' from the char.
			digit := int(char - '0')

			// Havn't found digit on line yet.
			if !foundDigit {
				firstDigit = digit
			}

			lastDigit = digit

			foundDigit = true
		}

    // Concatenate the two digits and add them to the sum when the end of the line is reached.
		if char == '\n' {
			sum += firstDigit*10 + lastDigit

			foundDigit = false
		}
	}

	return sum
}
