package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	elevationMap = make(map[string]string)
)

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)

	potentialTrailheads := [][]int{}

	for i, row := range strings.Split(input, "\n") {
		for j, elevation := range strings.Split(row, "") {
			if elevation == "0" {
				potentialTrailheads = append(potentialTrailheads, []int{i, j})
			}
			elevationMap[fmt.Sprintf("%d,%d", i, j)] = elevation
		}
	}

	totalScore := 0
	for _, potentialTrailhead := range potentialTrailheads {
		totalScore += traverse(potentialTrailhead[0], potentialTrailhead[1], 0)
	}

	fmt.Println(totalScore)
}

func traverse(x, y, targetElevation int) int {
	if elevationMap[fmt.Sprintf("%d,%d", x, y)] != fmt.Sprintf("%d", targetElevation) {
		return 0
	} else {
		if targetElevation == 9 {
			return 1
		}
	}
	return traverse(x+1, y, targetElevation+1) + traverse(x-1, y, targetElevation+1) + traverse(x, y+1, targetElevation+1) + traverse(x, y-1, targetElevation+1)
}
