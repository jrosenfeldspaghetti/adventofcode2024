package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var solvedStones = make(map[string][]string)

func main() {
	// sanityCheck("125 17", 25)

	input := "0 27 5409930 828979 4471 3 68524 170"
	currentStones := strings.Split(input, " ")
	numberOfBlinks := 75
	for i := 0; i < numberOfBlinks; i++ {
		currentStones = processStoneArray(currentStones)
	}
	fmt.Printf("Final stone count: %d", len(currentStones))
}

func sanityCheck(input string, blinks int) {
	currentStones := strings.Split(input, " ")
	for i := 0; i < blinks; i++ {
		if i < 7 {
			fmt.Printf("After %d blinks: %s\n", i, strings.Join(currentStones, " "))
		}
		currentStones = processStoneArray(currentStones)
	}

	if len(currentStones) != 55312 {
		log.Fatalf("Sanity check failed: %d!= 55312", len(currentStones))
	}
}

func processStoneArray(stones []string) []string {
	// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
	// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
	// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

	resultingStones := []string{}
	for _, stone := range stones {
		newStones := processStone(stone)
		resultingStones = append(resultingStones, newStones...)
	}

	return resultingStones
}

func processStone(stone string) []string {
	if precalculatedStones, arleadySolved := solvedStones[stone]; arleadySolved {
		return precalculatedStones
	}

	fmt.Println("This stone was not found yet", stone)
	if stone == "0" {
		solvedStones[stone] = []string{"1"}
		return []string{"1"}
	}

	if len(stone)%2 == 0 {
		leftHalf, _ := strconv.Atoi(stone[:len(stone)/2])
		rightHalf, _ := strconv.Atoi(stone[len(stone)/2:])
		solvedStones[stone] = []string{fmt.Sprintf("%d", leftHalf), fmt.Sprintf("%d", rightHalf)}
		return []string{fmt.Sprintf("%d", leftHalf), fmt.Sprintf("%d", rightHalf)}
	}

	stoneValue, _ := strconv.Atoi(stone)
	stoneValue = stoneValue * 2024
	return []string{fmt.Sprintf("%d", stoneValue)}
}
