package main

import (
	"fmt"
	"os"
	"strings"
)

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	wordSearch := string(fileContents)

	wordSearchMap := map[string]rune{}
	foundWords := 0
	// iterate left to right looking for 'X'. Upon each X, try each of the 8 directions to see if there is a 'MAS' in that direction
	for rowIndex, rows := range strings.Split(wordSearch, "\n") {
		for colIndex, letter := range rows {
			wordSearchMap[fmt.Sprintf("%d,%d", rowIndex, colIndex)] = letter
		}
	}
	for rowIndex, rows := range strings.Split(wordSearch, "\n") {
		for colIndex, letter := range rows {
			if letter != 'X' {
				continue
			}
			for _, direction := range directions {
				valid := true
				for i, desiredLetter := range "MAS" {
					mapKey := fmt.Sprintf("%d,%d", rowIndex+(direction[0]*(i+1)), colIndex+(direction[1]*(i+1)))
					if letterIsWhatWeWant(wordSearchMap, mapKey, desiredLetter) {
						continue
					} else {
						valid = false
						break
					}
				}
				if valid {
					foundWords++
				}
			}
		}
	}
	fmt.Println(foundWords)
}

func letterIsWhatWeWant(letterMap map[string]rune, mapKey string, desiredLetter rune) bool {
	return letterMap[mapKey] == desiredLetter
}
