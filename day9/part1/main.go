package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(inputBytes)
	file := true
	fileId := 0
	computerArray := []string{}
	// keep track of the .'s from left to right
	// keep track of the #'s from left to right
	// when to end?

	for _, block := range strings.Split(input, "") {
		blockInt, _ := strconv.Atoi(block)
		if file {
			for i := 0; i < blockInt; i++ {
				computerArray = append(computerArray, fmt.Sprintf("%d", fileId))
			}
			fileId++
			file = false
		} else {
			file = true
			for i := 0; i < blockInt; i++ {

				computerArray = append(computerArray, ".")
			}
		}
	}
	leftPointer := 0
	rightPointer := len(computerArray) - 1
	for leftPointer < rightPointer {
		for computerArray[leftPointer] != "." {
			if leftPointer > rightPointer {
				break
			}
			leftPointer++
		}
		for computerArray[rightPointer] == "." {
			if leftPointer > rightPointer {
				break
			}
			rightPointer--
		}
		temp := computerArray[rightPointer]
		computerArray[rightPointer] = computerArray[leftPointer]
		computerArray[leftPointer] = temp
		rightPointer--
		leftPointer++
	}

	// without this for loop, there is a situation where at the end of the
	// compressed memory, there is the order [. 5224]. By running the loop
	// again, this gets fixed. This is very expensive and I'm sure the logic
	// above could be corrected but it's nearly Christmas, spare me :(
	leftPointer = 0
	rightPointer = len(computerArray) - 1
	for leftPointer < rightPointer {
		for computerArray[leftPointer] != "." {
			if leftPointer > rightPointer {
				break
			}
			leftPointer++
		}
		for computerArray[rightPointer] == "." {
			if leftPointer > rightPointer {
				break
			}
			rightPointer--
		}
		temp := computerArray[rightPointer]
		computerArray[rightPointer] = computerArray[leftPointer]
		computerArray[leftPointer] = temp
		rightPointer--
		leftPointer++
	}

	total := 0
	for i, fileId := range computerArray {
		if fileId == "." {
			continue
		}
		fileIdInt, _ := strconv.Atoi(fileId)
		total += (i * fileIdInt)
	}

	fmt.Println(total)
}
