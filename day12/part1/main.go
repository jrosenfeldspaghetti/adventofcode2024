package main

import (
	"fmt"
	"log"
	"os"
)

func calculateCost(farm string) int {
	traverser := newRegionTraverser(farm)
	return traverser.calculateFenceCost()
}

func main() {
	if err := sanityChecks(); err != nil {
		log.Fatalf("Sanity checks failed: %v", err)
	}

	input, err := os.ReadFile("farmInput.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)

	fmt.Printf("The cost of the farm provided is %d in whatever currency elves use.\n", calculateCost(inputStr))
}

func sanityChecks() error {
	farm1 := `AAAA
BBCD
BBCC
EEEC`
	farm2 := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	farm3 := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	farm1Cost := calculateCost(farm1)
	if farm1Cost != 140 {
		return fmt.Errorf("farm 1 should have a cost of 140, but it has a cost of %d", farm1Cost)
	}
	farm2Cost := calculateCost(farm2)
	if farm2Cost != 772 {
		return fmt.Errorf("farm 2 should have a cost of 772, but it has a cost of %d", farm2Cost)
	}
	farm3Cost := calculateCost(farm3)
	if farm3Cost != 1930 {
		return fmt.Errorf("farm 3 should have a cost of 1930, but it has a cost of %d", farm3Cost)
	}
	return nil
}
