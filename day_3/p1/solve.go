// For more info see https://adventofcode.com/2024/day/3
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("resources/input.txt")

	// Print the error if something went wrong
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	result := 0

	buffer := bufio.NewReader(file)
	for {
		line, _, err := buffer.ReadLine()

		// Stop when the lines are completed
		if err != nil {
			break
		}

		// Parse the operations from the input string
		operations := scanInput(string(line))

		// Calculate the result of the operations
		for _, operation := range operations {
			println("Running ", string(operation))
			result += runOperation(string(operation))
		}

	}

	fmt.Println("Operations results: ", result)

}

// calculates the result of the mul operations
func runOperation(s string) int {
	// Remove the mul string
	s = strings.Replace(s, "mul(", "", -1)
	// Remove the trailing (
	s = s[:len(s)-1]

	// Extract the numbers
	numbers := strings.Split(s, ",")

	// Multiply them
	result := 1
	for _, n := range numbers {
		parsedN, _ := strconv.Atoi(n)
		result *= parsedN
	}

	return result
}

// Scan the string for the proper mul(\d,\d) substrings
func scanInput(s string) [][]byte {
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := r.FindAll([]byte(s), -1)

	return matches
}
