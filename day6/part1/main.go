package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	north          = []int{-1, 0}
	east           = []int{0, 1}
	south          = []int{1, 0}
	west           = []int{0, -1}
	directionOrder = [][]int{north, east, south, west}
	visitedSet     = make(map[string]bool)
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	// find guard position. iterate until guard position exceeds grid.
	rows := strings.Split(inputStr, "\n")
	roomGrid := [][]string{}
	roomMap := make(map[string]string)

	for _, row := range rows {
		roomGrid = append(roomGrid, strings.Split(row, ""))
	}

	startingPositionX, startingPositionY := 0, 0
	for row := range roomGrid {
		for col := range roomGrid[row] {
			if roomGrid[row][col] == "^" {
				startingPositionX = row
				startingPositionY = col
			}
			roomMap[fmt.Sprintf("%d,%d", row, col)] = roomGrid[row][col]
		}
	}

	traverse(startingPositionX, startingPositionY, roomMap, 0)
	fmt.Println(len(visitedSet))
}

func traverse(x, y int, roomMap map[string]string, directionIndex int) {
	mapKey := fmt.Sprintf("%d,%d", x, y)
	if _, inBounds := roomMap[mapKey]; !inBounds {
		return
	}
	// yay for sets. We might "register" ourselve at the same location multiple times, but it won't increase unique tile visits
	visitedSet[mapKey] = true
	if roomMap[fmt.Sprintf("%d,%d", x+directionOrder[directionIndex][0], y+directionOrder[directionIndex][1])] == "#" {
		// we turn, and try again.
		traverse(x, y, roomMap, (directionIndex+1)%4)
	} else {
		// we didn't bump into anything, so continue on
		traverse(x+directionOrder[directionIndex][0], y+directionOrder[directionIndex][1], roomMap, directionIndex)
	}
}
