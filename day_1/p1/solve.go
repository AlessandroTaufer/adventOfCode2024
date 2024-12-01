package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    // Let's read the input file
	file, err := os.Open("resources/input_1.txt")
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

	// Let's sort the two slices
	sort.Ints(v1)
	sort.Ints(v2)

    // Let's perform a sanity check
    if len(v1) != len(v2){
        fmt.Println("Something went wrong my dear friend")
        _ = errors.New("challenge input is not healthy, different len vectors detected")
    }

    // Let's now calculate the total distance
    sum := 0
    for i:=0; i<len(v1); i++ {
        sum += int(math.Abs(float64(v1[i] - v2[i])))
    }

    // Output the result, yay
    fmt.Println("The total distance is ", sum)

}
