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

	file := true
	input := string(inputBytes)
	fileId := 0
	memoryArray := []string{}
	fileSizeMap := make(map[int]int)
	// value is starting address of size
	fileAddressMap := make(map[int]int)

	for _, block := range strings.Split(input, "") {
		blockInt, _ := strconv.Atoi(block)
		if file {
			for i := 0; i < blockInt; i++ {
				memoryArray = append(memoryArray, fmt.Sprintf("%d", fileId))
			}
			fileSizeMap[fileId] = blockInt
			fileAddressMap[fileId] = len(memoryArray) - blockInt
			fileId++
			file = false
		} else {
			file = true
			for i := 0; i < blockInt; i++ {
				memoryArray = append(memoryArray, ".")
			}
		}
	}

	fileIdPointer := fileId - 1
	for fileIdPointer > 0 {
		for i := 0; i < len(memoryArray); i++ {
			stopLooking := false
			if i >= fileAddressMap[fileIdPointer] {
				stopLooking = true
				break
			}
			if memoryArray[i] == "." {
				freeBlockStartingAddress := i
				freeBlockSize := 1
				for i < len(memoryArray) && memoryArray[i] == "." {
					if freeBlockSize == fileSizeMap[fileIdPointer] {
						startingFileAddress := fileAddressMap[fileIdPointer]
						for moverIndex := 0; moverIndex < fileSizeMap[fileIdPointer]; moverIndex++ {
							memoryArray[startingFileAddress+moverIndex] = "."
							memoryArray[freeBlockStartingAddress+moverIndex] = fmt.Sprintf("%d", fileIdPointer)
						}
						stopLooking = true
						break
					}
					i++
					freeBlockSize++
				}
			}
			if stopLooking {
				break
			}
		}
		fileIdPointer--
	}

	total := 0
	for i, fileId := range memoryArray {
		if fileId == "." {
			continue
		}
		fileIdInt, _ := strconv.Atoi(fileId)
		total += (i * fileIdInt)
	}
	fmt.Println(total)
}
