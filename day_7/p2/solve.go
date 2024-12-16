// Additional info at https://adventofcode.com/2024/day/7#part2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("resources/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := bufio.NewReader(file)

	test_values_sum := 0
	for {
		line, _, err := buffer.ReadLine()

		// Stop when file is ended
		if err != nil {
			break
		}

		// Parse the input
		result, operands := parseCalibrationInput(string(line))

		// Calculate validity
		is_valid := validateEquation(result, operands)

		if is_valid {
			test_values_sum += result
			fmt.Println(result, ":", operands, " is valid")
		}
	}

	fmt.Println("The number of valid equations is:", test_values_sum)
}

func validateEquation(result int, operands []int) bool {
	// If the first operand it's greater than the result we can already considered the equation impossible
	if operands[0] > result {
		return false
	}

	// If the lenght of the operands if 1 it means the equation has been completely evaluated
	if len(operands) < 2 {
		return operands[0] == result
	}

	// Compute another step of the equation (for each possible operator)
	add_subset := []int{operands[0] + operands[1]}
	multiply_subset := []int{operands[0] * operands[1]}

	// Concatenate the two numbers
	concat_num := strconv.Itoa(operands[0]) + strconv.Itoa(operands[1])
	concat_num_int, _ := strconv.Atoi(concat_num)
	concatenation_subset := []int{concat_num_int}

	// Add the remaining operands if any
	if len(operands) >= 3 {
		add_subset = append(add_subset, operands[2:]...)
		multiply_subset = append(multiply_subset, operands[2:]...)
		concatenation_subset = append(concatenation_subset, operands[2:]...)
	}

	return validateEquation(result, add_subset) || validateEquation(result, multiply_subset) || validateEquation(result, concatenation_subset)
}

func parseCalibrationInput(s string) (int, []int) {
	splitted_by_column := strings.Split(s, ":")
	result, _ := strconv.Atoi(string(splitted_by_column[0]))

	raw_equation_numbers := strings.Split(splitted_by_column[1], " ")
	equation_numbers := []int{}
	for _, n := range raw_equation_numbers {
		// skip empty lines
		if n == "" {
			continue
		}

		int_n, _ := strconv.Atoi(n)
		equation_numbers = append(equation_numbers, int_n)
	}

	return result, equation_numbers
}
