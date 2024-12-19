package main

import (
	"fmt"
	"os"
	"strings"
)

var directions = map[string][]int{
	"topLeft":  {-1, -1, 1, 1},
	"topRight": {-1, 1, 1, -1},
}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	wordSearch := string(fileContents)

	wordSearchMap := map[string]rune{}
	foundWords := 0
	for rowIndex, rows := range strings.Split(wordSearch, "\n") {
		for colIndex, letter := range rows {
			wordSearchMap[fmt.Sprintf("%d,%d", rowIndex, colIndex)] = letter
		}
	}

	for rowIndex, rows := range strings.Split(wordSearch, "\n") {
		for colIndex, letter := range rows {
			if letter != 'A' {
				continue
			}
			oneHalfOfTheXExists := false
			for _, direction := range directions {
				locationOne := fmt.Sprintf("%d,%d", rowIndex+direction[0], colIndex+direction[1])
				locationTwo := fmt.Sprintf("%d,%d", rowIndex+direction[2], colIndex+direction[3])
				if makesMas(wordSearchMap, locationOne, locationTwo) {
					if oneHalfOfTheXExists {
						foundWords++
					}
					oneHalfOfTheXExists = true
				}
			}
		}
	}

	fmt.Println(foundWords)
}

func makesMas(letterMap map[string]rune, locationOne string, locationTwo string) bool {
	return letterMap[locationOne] == 'M' && letterMap[locationTwo] == 'S' || letterMap[locationOne] == 'S' && letterMap[locationTwo] == 'M'
}
