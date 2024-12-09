// For additional info see https://adventofcode.com/2024/day/5#part2
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Page struct {
	id           int
	dependencies []int
}

func main() {
	fmt.Println("Hello there")

	file, err := os.Open("resources/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	buffer := bufio.NewReader(file)

	pages_hashgraph := make(map[int]Page)

	// Load the page dependencies
	for {
		line, _, err := buffer.ReadLine()

		// Stop on error
		if err != nil {
			break
		}

		// Stop reading the dependency section if there is a \n
		if string(line) == "" {
			fmt.Println("Reading of the pages dependencies concluded")
			break
		}

		new_page := loadPageFromString(string(line))

		// Check if the page already exist in the dictionary
		pre_existing_page, already_present := pages_hashgraph[new_page.id]

		// If it doesn't exist add it to the dictionary
		if !already_present {
			pages_hashgraph[new_page.id] = new_page
		} else {
			// Otherwise merge it with the already present dependencies
			pre_existing_page.dependencies = append(pre_existing_page.dependencies, new_page.dependencies...)
			pages_hashgraph[new_page.id] = pre_existing_page
		}

	}

	fmt.Println("dependencies of 97", pages_hashgraph[97].dependencies)

	// Load the updates
	middle_page_sum := 0
	for {
		line, _, err := buffer.ReadLine()

		if err != nil {
			break
		}

		update := loadUpdateFromString(string(line))

		// If it is not already ordered, order it and add the middle page to the sum
		if !validateUpdate(update, pages_hashgraph) {
			sortedUpdate := sortUpdate(update, pages_hashgraph)
			middle_page_sum += sortedUpdate[len(update)/2]
			fmt.Println("Update:", update, " becomes:", sortedUpdate)
		}
	}

	fmt.Println("The correctly unordered updates sum is", middle_page_sum)
}

// From an input string generate the correspondin page obj
func loadPageFromString(s string) Page {
	numbers := strings.Split(s, "|")

	var p Page

	p.id, _ = strconv.Atoi(numbers[0])
	page_dependency, _ := strconv.Atoi(numbers[1])

	p.dependencies = append(p.dependencies, page_dependency)

	return p
}

// Load from the input string the slice of integer corresponding to the pages id
func loadUpdateFromString(s string) []int {
	splitted_s := strings.Split(s, ",")
	var update []int

	for _, n := range splitted_s {
		n_int, _ := strconv.Atoi(n)
		update = append(update, n_int)
	}

	return update
}

// Validates a whole update
func validateUpdate(update_ids []int, pages_hashgraph map[int]Page) bool {
	// Sequence of pages that have already been validated
	// The first one can not be rejected so let's directly add it to be efficient
	validated_sequence := []int{update_ids[0]}

	// For each candidate page check if it's dependencies conflict with the already validated sequence
	for _, candidate_page := range update_ids[1:] {
		for _, validated_page := range validated_sequence {
			is_valid := inspectDependencies(candidate_page, validated_page, pages_hashgraph)

			// Stop as soon as you find an invalid page
			if !is_valid {
				return false
			}
		}

		validated_sequence = append(validated_sequence, candidate_page)
	}

	return true
}

// Given an unsorted update, returns a sorted one
func sortUpdate(update_ids []int, pages_hashgraph map[int]Page) []int {
	sorted_ids := []int{update_ids[0]}

	// For each candidate page check if it's dependencies conflict with the already validated sequence
	for _, candidate_page := range update_ids[1:] {
		is_ordered := true
		for i, validated_page := range sorted_ids {
			is_ordered = inspectDependencies(candidate_page, validated_page, pages_hashgraph)

			// If you find an invalid page, you can assume that the candidate page should be inserted before it
			if !is_ordered {
				sorted_ids = slices.Insert(sorted_ids, i, candidate_page)
				break
			}
		}
		// If it was already ordered, put it at the end of the update
		if is_ordered {
			sorted_ids = append(sorted_ids, candidate_page)
		}
	}

	return sorted_ids
}

// Inspect the dependencies of the candidate page ensuring the already validated page is not there
func inspectDependencies(candidate_page int, validated_page int, pages_hashgraph map[int]Page) bool {
	// If the candidate page is present in the dependencies, then the order it's not valid
	return !slices.Contains(pages_hashgraph[candidate_page].dependencies, validated_page)
}

// Returns true if the num
