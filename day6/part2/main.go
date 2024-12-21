package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	startingPositionX = 0
	startingPositionY = 0
	north             = []int{-1, 0}
	east              = []int{0, 1}
	south             = []int{1, 0}
	west              = []int{0, -1}
	directionOrder    = [][]int{north, east, south, west}
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	// brute force! What if we just tried putting an obstacle in every position that isn't the starting position and see if we can
	// detect a loop or not
	// find guard position. iterate until guard position exceeds grid.
	rows := strings.Split(inputStr, "\n")
	roomGrid := [][]string{}
	roomMap := make(map[string]string)

	for _, row := range rows {
		roomGrid = append(roomGrid, strings.Split(row, ""))
	}

	for row := range roomGrid {
		for col := range roomGrid[row] {
			if roomGrid[row][col] == "^" {
				startingPositionX = row
				startingPositionY = col
			}
			roomMap[fmt.Sprintf("%d,%d", row, col)] = roomGrid[row][col]
		}
	}
	loopLocations := 0
	for row := range roomGrid {
		for col := range roomGrid[row] {
			if roomGrid[row][col] == "." {
				testMap := make(map[string]string)
				for k, v := range roomMap {
					testMap[k] = v
				}
				testMap[fmt.Sprintf("%d,%d", row, col)] = "#"
				if hasLoop(testMap) {
					loopLocations++
				}
			}
		}
	}

	fmt.Println("loop locations:", loopLocations)
}

func hasLoop(roomMap map[string]string) bool {
	visitedSet := make(map[string]bool)
	x, y := startingPositionX, startingPositionY
	directionIndex := 0
	_, inRange := roomMap[fmt.Sprintf("%d,%d", startingPositionX, startingPositionY)]
	for inRange {
		// if we have been here already in the same direction we're traveling, let's assume that's a loop
		if alreadyVisited := visitedSet[fmt.Sprintf("%d,%d,%d", x, y, directionIndex)]; alreadyVisited {
			return true
		}
		visitedSet[fmt.Sprintf("%d,%d,%d", x, y, directionIndex)] = true
		if roomMap[fmt.Sprintf("%d,%d", x+directionOrder[directionIndex][0], y+directionOrder[directionIndex][1])] == "#" {
			directionIndex = (directionIndex + 1) % 4
			continue
		}
		x = x + directionOrder[directionIndex][0]
		y = y + directionOrder[directionIndex][1]
		_, inRange = roomMap[fmt.Sprintf("%d,%d", x, y)]
	}
	return false
}
