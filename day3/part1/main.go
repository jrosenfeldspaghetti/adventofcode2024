package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	mulRegexPattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`
	digitPattern := `[0-9]{1,3}`
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	command := string(content)

	compiledMulPattern := regexp.MustCompile(mulRegexPattern)
	compiledDigitPattern := regexp.MustCompile(digitPattern)

	total := getTotalFromString(compiledMulPattern, command, compiledDigitPattern)

	fmt.Printf("Command output is %d\n", total)
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
