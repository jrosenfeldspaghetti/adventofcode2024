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
		foundPoints := make(map[string]bool)
		traverse(potentialTrailhead[0], potentialTrailhead[1], 0, foundPoints)
		totalScore += len(foundPoints)
	}

	fmt.Println(totalScore)
}

func traverse(x, y, targetElevation int, foundPoints map[string]bool) {
	if elevationMap[fmt.Sprintf("%d,%d", x, y)] != fmt.Sprintf("%d", targetElevation) {
		return
	} else {
		if targetElevation == 9 {
			foundPoints[fmt.Sprintf("%d,%d", x, y)] = true
			return
		}
	}
	traverse(x+1, y, targetElevation+1, foundPoints)
	traverse(x-1, y, targetElevation+1, foundPoints)
	traverse(x, y+1, targetElevation+1, foundPoints)
	traverse(x, y-1, targetElevation+1, foundPoints)
}
