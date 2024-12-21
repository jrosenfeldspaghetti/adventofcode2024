package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	multiply    = "multiply"
	add         = "add"
	concatenate = "concatenate"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)
	total := 0
	for _, calculation := range strings.Split(inputStr, "\n") {
		splitOnColon := strings.Split(calculation, ":")
		goal, _ := strconv.Atoi(splitOnColon[0])
		numbersToCalc := []int{}
		for _, num := range strings.Split(strings.TrimSpace(splitOnColon[1]), " ") {
			numAsInt, _ := strconv.Atoi(num)
			numbersToCalc = append(numbersToCalc, numAsInt)
		}
		if calculateOption(0, goal, numbersToCalc, add) ||
			calculateOption(0, goal, numbersToCalc, multiply) ||
			calculateOption(0, goal, numbersToCalc, concatenate) {
			total += goal
		}

	}
	fmt.Println(total)
}

func calculateOption(current, goal int, vals []int, operator string) bool {
	if current == goal && len(vals) == 0 {
		return true
	}
	// short circuit some calculations
	if current > goal {
		return false
	}

	if operator == multiply {
		newCurrent := current * vals[0]
		if len(vals) == 1 {
			return newCurrent == goal
		}
		if calculateOption(newCurrent, goal, vals[1:], multiply) ||
			calculateOption(newCurrent, goal, vals[1:], add) ||
			calculateOption(newCurrent, goal, vals[1:], concatenate) {
			return true
		}
	}
	if operator == add {
		newCurrent := current + vals[0]
		if len(vals) == 1 {
			return newCurrent == goal
		}
		if calculateOption(newCurrent, goal, vals[1:], multiply) ||
			calculateOption(newCurrent, goal, vals[1:], add) ||
			calculateOption(newCurrent, goal, vals[1:], concatenate) {
			return true
		}
	}

	if operator == concatenate {
		newCurrent, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, vals[0]))
		if len(vals) == 1 {
			return newCurrent == goal
		}
		if calculateOption(newCurrent, goal, vals[1:], multiply) ||
			calculateOption(newCurrent, goal, vals[1:], add) ||
			calculateOption(newCurrent, goal, vals[1:], concatenate) {
			return true
		}
	}

	return false
}
