package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sanityCheck()
	leftList := []int{}
	rightList := []int{}

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, " ")
		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[len(numbers)-1])
		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}
	fmt.Printf("Total similarityScore: %d\n", calculateSimilarityScore(leftList, rightList))
}

func sanityCheck() {
	leftList := []int{3, 4, 2, 1, 3, 3}
	rightList := []int{4, 3, 5, 3, 9, 3}

	if calculateSimilarityScore(leftList, rightList) != 31 {
		panic("from test data, similarity score should be 31")
	}
}

func calculateSimilarityScore(leftList, rightList []int) int {
	ocurrenceMap := make(map[int]int)

	for _, rightNumber := range rightList {
		ocurrenceMap[rightNumber] += 1
	}

	similarityScore := 0
	for _, leftNumber := range leftList {
		similarityScore += ocurrenceMap[leftNumber] * leftNumber
	}

	return similarityScore
}
