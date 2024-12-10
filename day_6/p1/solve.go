// For additional info see https://adventofcode.com/2024/day/6
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("resources/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	buffer := bufio.NewReader(file)

	var terrain [][]rune

	// Let's load the terrain from file
	for {
		line, _, err := buffer.ReadLine()

		if err != nil {
			break
		}

		var characters []rune
		for _, c := range string(line) {
			characters = append(characters, c)
		}

		terrain = append(terrain, characters)
	}

	printTerrain(terrain)
	guardX, guardY := detectGuardCoordinates(terrain)
	distinctPositionsCounter := moveGuard(0, guardX, guardY, 0, terrain)
	fmt.Println("Unique positions visited by the guard:", distinctPositionsCounter)
}

// Return the current position of the guard in the terrain
func detectGuardCoordinates(terrain [][]rune) (int, int) {
	for i := 0; i < len(terrain); i++ {
		for j := 0; j < len(terrain[i]); j++ {
			if terrain[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func moveGuard(direction int, cursorX int, cursorY int, distinctPositionVisited int, terrain [][]rune) int {
	// calculate the modifiers that will affect the movement of the guard
	var movementModifiers []int

	// up
	if direction == 0 {
		movementModifiers = []int{-1, 0}

		// right
	} else if direction == 1 {
		movementModifiers = []int{0, 1}
		// down
	} else if direction == 2 {
		movementModifiers = []int{1, 0}
		// left
	} else {
		movementModifiers = []int{0, -1}
	}

	// If the terrain is not already visited increment the counter of position visited
	if terrain[cursorX][cursorY] != 'X' {
		distinctPositionVisited++

		// Mark the current position as visited in the map
		terrain[cursorX][cursorY] = 'X'
	}

	// Check if the next direction will be outside the map, if so stop the map
	if cursorX+movementModifiers[0] < 0 || cursorX+movementModifiers[0] >= len(terrain[0]) ||
		cursorY+movementModifiers[1] < 0 || cursorY+movementModifiers[1] >= len(terrain[1]) {
		return distinctPositionVisited
	}

	// check if the next position is obstructed, if so rotate it and iterate the function
	if terrain[cursorX+movementModifiers[0]][cursorY+movementModifiers[1]] == '#' {
		newDirection := (direction + 1) % 4
		return moveGuard(newDirection, cursorX, cursorY, distinctPositionVisited, terrain)
	}

	//printTerrain(terrain)
	return moveGuard(direction, cursorX+movementModifiers[0], cursorY+movementModifiers[1], distinctPositionVisited, terrain)
}

// Print the current content of the terrain
func printTerrain(terrain [][]rune) {
	fmt.Print("\n<============>\n")
	for _, line := range terrain {
		fmt.Println(string(line))
	}

}
