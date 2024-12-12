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
	farm3 := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	farm4 := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	farm5 := `RRRRIICCFF
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
	if farm1Cost != 80 {
		return fmt.Errorf("farm 1 should have a cost of 80, but it has a cost of %d", farm1Cost)
	}
	farm2Cost := calculateCost(farm2)
	if farm2Cost != 436 {
		return fmt.Errorf("farm 2 should have a cost of 436, but it has a cost of %d", farm2Cost)
	}
	farm3Cost := calculateCost(farm3)
	if farm3Cost != 236 {
		return fmt.Errorf("farm 3 should have a cost of 236, but it has a cost of %d", farm3Cost)
	}
	farm4Cost := calculateCost(farm4)
	if farm4Cost != 368 {
		return fmt.Errorf("farm 4 should have a cost of 368, but it has a cost of %d", farm4Cost)
	}
	farm5Cost := calculateCost(farm5)
	if farm5Cost != 1206 {
		return fmt.Errorf("farm 5 should have a cost of 1206, but it has a cost of %d", farm5Cost)
	}
	return nil
}
