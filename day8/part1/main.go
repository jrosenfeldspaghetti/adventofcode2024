package main

import (
	"fmt"
	"os"
	"strings"
)

type antenna struct {
	x int
	y int
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(inputBytes)
	antennaLocations := make(map[string][]antenna)
	inputAsMap := make(map[string]string)

	for row, inputRow := range strings.Split(input, "\n") {
		for col, point := range strings.Split(inputRow, "") {
			inputAsMap[fmt.Sprintf("%d,%d", row, col)] = point
			if point != "." {
				antennaLocations[point] = append(antennaLocations[point], antenna{x: row, y: col})
			}
		}
	}

	// for each antenna, for each pairing, calculate their two antinodes and check that they are within bounds
	antinodes := make(map[string]bool)
	for antennaKey := range antennaLocations {
		for a1Index, a1 := range antennaLocations[antennaKey] {
			a2Index := a1Index + 1
			for a2Index < len(antennaLocations[antennaKey]) {
				a2 := antennaLocations[antennaKey][a2Index]
				rowDistance := abs(a1.x - a2.x)
				colDistance := abs(a1.y - a2.y)
				antinode1X := a1.x - rowDistance
				antinode2X := a2.x + rowDistance
				if a1.x > a2.x {
					antinode1X = a1.x + rowDistance
					antinode2X = a2.x - rowDistance
				}
				antinode1Y := a1.y - colDistance
				antinode2Y := a2.y + colDistance
				if a1.y > a2.y {
					antinode1Y = a1.y + colDistance
					antinode2Y = a2.y - colDistance

				}

				antinode1Key := fmt.Sprintf("%d,%d", antinode1X, antinode1Y)
				antinode2Key := fmt.Sprintf("%d,%d", antinode2X, antinode2Y)

				if _, inBounds := inputAsMap[antinode1Key]; inBounds {
					antinodes[antinode1Key] = true
				}
				if _, inBounds := inputAsMap[antinode2Key]; inBounds {
					antinodes[antinode2Key] = true
				}

				a2Index++
			}
		}
	}
	fmt.Println(len(antinodes))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
