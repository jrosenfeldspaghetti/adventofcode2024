package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var relationMap = make(map[string][]string)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)
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

	validUpdates := [][]string{}
	for _, update := range pagesToPrint {
		fixedList, neededFixing := fixList(strings.Split(update, ","), false)
		if neededFixing {
			validUpdates = append(validUpdates, fixedList)
		}
	}

	total := 0
	for _, update := range validUpdates {
		val, _ := strconv.Atoi(update[(len(update)-1)/2])
		total += val
	}
	fmt.Println(total)
}

func fixList(updates []string, neededFixing bool) ([]string, bool) {
	indexThatLetterMustBeInFrontOf := make(map[string]int)
	for pageIndex, page := range updates {
		// first, check the relation map and mark all numbers that have to be before this point in the update list
		if restrictions, hasRestrictions := relationMap[page]; hasRestrictions {
			for _, restriction := range restrictions {
				// if this number already needs to be earlier in the list, we can skip marking it.
				if _, seenPageNumberBefore := indexThatLetterMustBeInFrontOf[restriction]; !seenPageNumberBefore {
					indexThatLetterMustBeInFrontOf[restriction] = pageIndex
				}
			}
		}
		// then, check the marked numbers. Is this page in that list? If so, it needs to be moved before that pivot point
		if pivotIndex, collision := indexThatLetterMustBeInFrontOf[page]; collision {
			nextList := []string{}
			nextList = append(nextList, updates...)
			t := nextList[pivotIndex]
			nextList[pivotIndex] = nextList[pageIndex]
			nextList[pageIndex] = t
			return fixList(nextList, true)
		}

	}
	return updates, neededFixing
}
