// For more info see https://adventofcode.com/2024/day/1
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Let's read the input file
	file, err := os.Open("resources/input.txt")
	if err != nil {
		fmt.Println("failed while opening the file", err)
	}
	defer file.Close()

	v1 := []int{}
	v2 := []int{}

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		if len(line) > 0 {
			str_line := string(line)
			e1 := strings.Split(str_line, " ")

			// Read the element belonging to the first list
			first_element, err := strconv.Atoi(e1[0])
			if err != nil {
				fmt.Println("Error while converting line to int", err)
			}
			v1 = append(v1, first_element)

			// Read the element belonging to the second list
			second_element, err := strconv.Atoi(e1[3])
			if err != nil {
				fmt.Println("Error while converting line to int", err)
			}
			v2 = append(v2, second_element)
		}
	}

	// Let's perform a sanity check
	if len(v1) != len(v2) {
		fmt.Println("Something went wrong my dear friend")
		_ = errors.New("challenge input is not healthy, different len vectors detected")
	}

	// We will use some dynamic programming here, otherwise my pc is gonna melt
	// to do so, we will keep map that tracks all the already calculated occurrences of each id, for each vector
	occurrences_v1 := countRepeatedItems(v1)
	occurrences_v2 := countRepeatedItems(v2)

	// We can now calculate the similarity by just multpling the id with the occurences in the two vectors
	similarity := 0
	for id, occurences_in_v1 := range occurrences_v1 {
		similarity += id * occurences_in_v1 * occurrences_v2[id]
	}

	// Output the result, yay
	fmt.Println("The similarity score is ", similarity)

}

// Return a map containing the occurences of each unique element
func countRepeatedItems(v []int) map[int]int {
	occurrences := make(map[int]int)
	for i := 0; i < len(v); i++ {
		items_already_counted, already_counted := occurrences[v[i]]

		// Init the map to 1 if not already initialized
		if !already_counted {
			occurrences[v[i]] = 1
			continue
		} else {
			// Increment it otherwise
			occurrences[v[i]] = items_already_counted + 1
		}
	}
	return occurrences
}
