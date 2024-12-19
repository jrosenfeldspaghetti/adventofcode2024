package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// first find all dos and dont's. find the index in which they end
	// build from there valid substrings to do the matchings for, calculate those, and add them to the total

	doRegexPattern := `do\(\)`
	dontRegexPattern := `don't()`
	mulRegexPattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`
	digitPattern := `[0-9]{1,3}`

	compiledDoRegex := regexp.MustCompile(doRegexPattern)
	compiledDontRegex := regexp.MustCompile(dontRegexPattern)
	compiledMulRegex := regexp.MustCompile(mulRegexPattern)
	compiledDigitRegex := regexp.MustCompile(digitPattern)

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	command := string(content)

	validSlices := []string{}
	doIndices := compiledDoRegex.FindAllStringIndex(command, -1)
	dontIndices := compiledDontRegex.FindAllStringIndex(command, -1)
	fmt.Println(dontIndices)

	// find the next dont block. Build first slice from that range.
	// find the next do block that enables section again. jump pointer there. repeat
	pointer := 0
	for pointer < len(command) {
		// basically, at the start of each iteration, we are in a "do" block.
		foundAnotherDontBlock := false
		for _, dontIndex := range dontIndices {
			if dontIndex[0] < pointer {
				continue
			}
			validSlices = append(validSlices, command[pointer:dontIndex[0]])
			pointer = dontIndex[1]
			foundAnotherDontBlock = true
			break
		}
		if !foundAnotherDontBlock {
			validSlices = append(validSlices, command[pointer:])
			break
		}
		foundAnotherDoBlock := false
		for _, doIndex := range doIndices {
			// the next do that is beyond where we are pointing now is where we start counting
			// things again.
			if doIndex[0] >= pointer {
				pointer = doIndex[1]
				foundAnotherDoBlock = true
				break
			}
		}
		if !foundAnotherDoBlock {
			// the rest of the string is in a don't block and can't be counted
			break
		}
	}

	grandTotal := 0
	for _, validSlice := range validSlices {
		grandTotal += getTotalFromString(compiledMulRegex, validSlice, compiledDigitRegex)
	}

	output := strings.Join(validSlices, "\n")
	err = os.WriteFile("validSlices.txt", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(grandTotal)
}

func getTotalFromString(compiledMulPattern *regexp.Regexp, command string, compiledDigitPattern *regexp.Regexp) int {
	total := 0
	if allMultiplyCommands := compiledMulPattern.FindAllStringSubmatch(command, -1); allMultiplyCommands != nil {
		for _, multiplyCommand := range allMultiplyCommands {
			digits := compiledDigitPattern.FindAllStringSubmatch(multiplyCommand[0], -1)
			firstVal, _ := strconv.Atoi(digits[0][0])
			secondVal, _ := strconv.Atoi(digits[1][0])
			total += firstVal * secondVal
		}
	} else {
		log.Fatalf("Unable to parse command input")
	}
	return total
}
