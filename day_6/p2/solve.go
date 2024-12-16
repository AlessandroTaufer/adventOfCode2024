// For additional info see https://adventofcode.com/2024/day/6#part2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	guardX, guardY := detectGuardCoordinates(terrain)
	loops_counter := generateObstructions(guardX, guardY, terrain)
	fmt.Println("Number of loops that can be created:", loops_counter)
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

func generateObstructions(guardX int, guardY int, terrain [][]rune) int {
	loops_counter := 0
	for i := 0; i < len(terrain); i++ {
		for j := 0; j < len(terrain[i]); j++ {
			terrain = resetTerrain(guardX, guardY, terrain)
			// If it's empty place an obstruction there
			if terrain[i][j] == '.' {
				terrain[i][j] = '#'
				// If adding the obstruction caused a loop increment the counter
				if moveGuard(0, guardX, guardY, terrain) {
					loops_counter++
				}

				// Restore the original terrain
				terrain[i][j] = '.'
			}
		}
	}
	return loops_counter
}

func resetTerrain(guardX int, guardY int, terrain [][]rune) [][]rune {
	for i := 0; i < len(terrain); i++ {
		for j := 0; j < len(terrain[i]); j++ {
			if i == guardX && j == guardY {
				terrain[i][j] = '^'
			}

			// If there is a guard movement character replace it with an empty one 
			if isRuneAMovementDirection(terrain[i][j]) {
				terrain[i][j] = '.'
			}
		}
	}
	return terrain
}


func isRuneAMovementDirection(r rune) bool{
	movementDirections := "↑→↓←"
	
	// If there is a guard movement character replace it with an empty one 
	return strings.Contains(movementDirections, string(r))
}

func moveGuard(direction int, cursorX int, cursorY int, terrain [][]rune) bool {
	// calculate the modifiers that will affect the movement of the guard
	var movementModifiers []int
	var movementDirection rune

	// up
	if direction == 0 {
		movementModifiers = []int{-1, 0}
		movementDirection = '↑'
		// right
	} else if direction == 1 {
		movementModifiers = []int{0, 1}
		movementDirection = '→'
		// down
	} else if direction == 2 {
		movementModifiers = []int{1, 0}
		movementDirection = '↓'
		// left
	} else {
		movementModifiers = []int{0, -1}
		movementDirection = '←'
	}

	// If the terrain is not already visited mark it as visited with the current direction
	if terrain[cursorX][cursorY] == '.' || terrain[cursorX][cursorY] == '^'{
		terrain[cursorX][cursorY] = movementDirection
	
	// If it has been visited in the same direction it mean there is a loop!
	}else if terrain[cursorX][cursorY] == movementDirection{
		fmt.Println("This is a loop!")
		printTerrain(terrain)
		return true
	}

	// Check if the next direction will be outside the map, if so stop the search
	if cursorX+movementModifiers[0] < 0 || cursorX+movementModifiers[0] >= len(terrain[0]) ||
		cursorY+movementModifiers[1] < 0 || cursorY+movementModifiers[1] >= len(terrain[1]) {
		return false
	}

	// check if the next position is obstructed, if so rotate it and iterate the function
	if terrain[cursorX+movementModifiers[0]][cursorY+movementModifiers[1]] == '#' {
		newDirection := (direction + 1) % 4
		return moveGuard(newDirection, cursorX, cursorY, terrain)
	}

	return moveGuard(direction, cursorX+movementModifiers[0], cursorY+movementModifiers[1], terrain)
}

// Print the current content of the terrain
func printTerrain(terrain [][]rune) {
	fmt.Print("\n<============>\n")
	for _, line := range terrain {
		fmt.Println(string(line))
	}

}
