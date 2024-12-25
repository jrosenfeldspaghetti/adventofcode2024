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
				dx1 := -rowDistance
				dx2 := rowDistance
				if a1.x > a2.x {
					dx1 = rowDistance
					dx2 = -rowDistance
				}
				dy1 := -colDistance
				dy2 := colDistance
				if a1.y > a2.y {
					dy1 = colDistance
					dy2 = -colDistance
				}

				x := a1.x
				y := a1.y
				for {
					if _, stillOnGrid := inputAsMap[fmt.Sprintf("%d,%d", x+dx1, y+dy1)]; !stillOnGrid {
						break
					}
					antinodeKey := fmt.Sprintf("%d,%d", x+dx1, y+dy1)
					if _, inBounds := inputAsMap[antinodeKey]; inBounds {
						antinodes[antinodeKey] = true
					}
					x = x + dx1
					y = y + dy1
				}

				x = a2.x
				y = a2.y
				for {
					if _, stillOnGrid := inputAsMap[fmt.Sprintf("%d,%d", x+dx2, y+dy2)]; !stillOnGrid {
						break
					}
					antinodeKey := fmt.Sprintf("%d,%d", x+dx2, y+dy2)
					if _, inBounds := inputAsMap[antinodeKey]; inBounds {
						antinodes[antinodeKey] = true
					}
					x = x + dx2
					y = y + dy2
				}

				x = a1.x
				y = a1.y
				for {
					if _, stillOnGrid := inputAsMap[fmt.Sprintf("%d,%d", x+dx2, y+dy2)]; !stillOnGrid {
						break
					}
					antinodeKey := fmt.Sprintf("%d,%d", x+dx2, y+dy2)
					if _, inBounds := inputAsMap[antinodeKey]; inBounds {
						antinodes[antinodeKey] = true
					}
					x = x + dx2
					y = y + dy2
				}

				x = a2.x
				y = a2.y
				for {
					if _, stillOnGrid := inputAsMap[fmt.Sprintf("%d,%d", x+dx1, y+dy1)]; !stillOnGrid {
						break
					}
					antinodeKey := fmt.Sprintf("%d,%d", x+dx1, y+dy1)
					if _, inBounds := inputAsMap[antinodeKey]; inBounds {
						antinodes[antinodeKey] = true
					}
					x = x + dx1
					y = y + dy1
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
