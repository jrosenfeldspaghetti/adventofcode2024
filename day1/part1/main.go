package main

import (
	"fmt"
	"os"
	"sort"
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

	fmt.Printf("Total distance: %d\n", calculateDistance(leftList, rightList))
}

func sanityCheck() {
	leftList := []int{3, 4, 2, 1, 3, 3}
	rightList := []int{4, 3, 5, 3, 9, 3}

	distance := calculateDistance(leftList, rightList)
	if distance != 11 {
		panic("from test data, distance should be 11")
	}
}

func calculateDistance(leftList, rightList []int) int {
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		if leftList[i] != rightList[i] {
			totalDistance += abs(leftList[i] - rightList[i])
		}
	}
	return totalDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
