// For more info see https://adventofcode.com/2024/day/2
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Let's open the file
	file, err := os.Open("resources/input.txt")
	if err != nil {
		fmt.Println("Error while opening input file", err)
	}
	defer file.Close()

	reports := [][]int{}
	reports_len := 0

	// And let's read the input
	buffer := bufio.NewReader(file)
	for {
		line, _, err := buffer.ReadLine()

		// Stop scanning if the end of line is reached
		if err != nil {
			break
		}

		splitted_values := strings.Split(string(line), " ")

		// Load the integers into the matrix
		reports = append(reports, []int{})
		for _, e := range splitted_values {
			// Don't append empty strings
			if e == "" {
				continue
			}
			num, _ := strconv.Atoi(e)
			reports[reports_len] = append(reports[reports_len], num)
		}

		reports_len++
	}

	safeReports := 0

	// Check the validity of each report
	for i := 0; i < len(reports); i++ {
		// fmt.Println(reports[i], " -- len: ", len(reports[i]))

		if checkTrend(reports[i]) {
			safeReports++
			// fmt.Println("Valid report!")
		}
	}

	fmt.Println("The valid reports are", safeReports)

}

// Return true if n1 and n2 are in ascending order, use reverse to invert it
func compareAscending(n1 int, n2 int, reverse bool) bool {
	if !reverse {
		return n1 < n2
	} else {
		return n1 > n2
	}
}

// Check if delta between is between 1 and 3 (included)
func checkLevelsDiffer(n1 int, n2 int) bool {
	delta := int(math.Abs(float64(n1 - n2)))
	return delta <= 3 && delta >= 1

}

func checkTrend(v []int) bool {
	isAscending := v[0] < v[1]
	for i := 0; i < len(v)-1; i++ {
		// Fail if ascending order is not respected
		if !compareAscending(v[i], v[i+1], !isAscending) {
			return false
		}
		if !checkLevelsDiffer(v[i], v[i+1]) {
			return false
		}
	}

	return true
}
