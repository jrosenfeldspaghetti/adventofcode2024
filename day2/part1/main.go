package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	testCase()
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	reports := strings.Split(inputString, "\n")
	safeReports := getSafeReportCount(reports)
	fmt.Printf("There are %d safe reports.\n", safeReports)
}

func testCase() {
	testInput := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	reports := strings.Split(testInput, "\n")
	safeReports := getSafeReportCount(reports)
	if safeReports != 2 {
		panic(fmt.Sprintf("test failed, as safeReports was not 2 but %d\n ", safeReports))
	}
}

func getSafeReportCount(reports []string) int {
	safeReports := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")
		if isIncreasing(levels) || isDecreasing(levels) {
			safeReports++
		}
	}
	return safeReports
}

func isIncreasing(sequenceStr []string) bool {
	firstVal, _ := strconv.Atoi(sequenceStr[0])
	previousVal := firstVal
	for i, element := range sequenceStr {
		if i == 0 {
			continue
		}
		elementVal, _ := strconv.Atoi(element)
		if elementVal > previousVal && differenceIsAcceptable(previousVal, elementVal) {
			previousVal = elementVal
			continue
		}
		return false
	}
	return true
}
func isDecreasing(sequenceStr []string) bool {
	firstVal, _ := strconv.Atoi(sequenceStr[0])
	previousVal := firstVal
	for i, element := range sequenceStr {
		if i == 0 {
			continue
		}
		elementVal, _ := strconv.Atoi(element)
		if elementVal < previousVal && differenceIsAcceptable(previousVal, elementVal) {
			previousVal = elementVal
			continue
		}
		return false
	}
	return true
}

func differenceIsAcceptable(int1, int2 int) bool {
	return abs(int1-int2) > 0 && abs(int1-int2) <= 3
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
