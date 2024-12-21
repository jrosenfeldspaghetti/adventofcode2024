package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)
	// the key cannot show up before the value
	relationMap := make(map[string][]string)
	pagesToPrint := []string{}

	for _, line := range strings.Split(inputStr, "\n") {
		if line == "" {
			continue
		}
		if numbers := strings.Split(line, "|"); len(numbers) > 1 {
			if _, ok := relationMap[numbers[1]]; !ok {
				relationMap[numbers[1]] = []string{}
			}
			relationMap[numbers[1]] = append(relationMap[numbers[1]], numbers[0])
			continue
		}
		pagesToPrint = append(pagesToPrint, line)
	}

	validUpdates := []string{}
	for _, update := range pagesToPrint {
		pagesThatCantBeVisited := make(map[string]bool)
		valid := true
		for _, page := range strings.Split(update, ",") {
			if _, inMap := pagesThatCantBeVisited[page]; inMap {
				valid = false
				break
			}
			if _, hasRestrictions := relationMap[page]; hasRestrictions {
				for _, restriction := range relationMap[page] {
					pagesThatCantBeVisited[restriction] = true
				}
			}
		}
		if valid {
			validUpdates = append(validUpdates, update)
		}
	}
	total := 0
	for _, update := range validUpdates {
		updateSlice := strings.Split(update, ",")
		val, _ := strconv.Atoi(updateSlice[(len(updateSlice)-1)/2])
		total += val
	}
	fmt.Println(total)
}
