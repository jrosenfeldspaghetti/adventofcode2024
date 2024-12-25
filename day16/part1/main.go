package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

const (
	east  = 0
	south = 1
	west  = 2
	north = 3
)

var directionVectors = [][]int{
	{0, 1},  // east
	{1, 0},  // south
	{0, -1}, // west
	{-1, 0}, // north
}

func main() {
	sanityCheck()

	// input, err := os.ReadFile("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// inputString := string(input)
	// mazeMatrix := buildMazeMatrix(inputString)
	// scoreMatrix := buildScoreMatrix(mazeMatrix)
	// startingX, startingY := getStartingPoint(mazeMatrix)
	// traverse(mazeMatrix, startingX, startingY, scoreMatrix, buildVisitedMatrix(mazeMatrix), 0, east)

	// fmt.Printf("The cheapest route is %d.\n", findCheapestRoute(mazeMatrix, scoreMatrix))

}

func buildScoreMatrix(mazeMatrix [][]string) [][]int {
	scoreMatrix := make([][]int, len(mazeMatrix))
	for i := range scoreMatrix {
		scoreMatrix[i] = make([]int, len(mazeMatrix[i]))
		for j := range scoreMatrix[i] {
			scoreMatrix[i][j] = int(^uint(0) >> 1)
		}
	}
	return scoreMatrix
}

func buildVisitedMatrix(mazeMatrix [][]string) [][]bool {
	visitedMatrix := make([][]bool, len(mazeMatrix))
	for i := range visitedMatrix {
		visitedMatrix[i] = make([]bool, len(mazeMatrix[i]))
		for j := range visitedMatrix[i] {
			visitedMatrix[i][j] = false
		}
	}
	return visitedMatrix
}

type h func(int, int, int) int

func heuristicFunc(currentDirectionIndex int, diffX, diffY int) int {
	directionX := directionVectors[currentDirectionIndex][0]
	directionY := directionVectors[currentDirectionIndex][1]
	if abs(diffX-directionX)+abs(diffY-directionY) == 0 {
		return 1
	}
	return 1001
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}

func aStarTraversal(mapMatrix [][]string, startingX, startingY int) int {
	openSet := PriorityQueue{&Node{x: startingX, y: startingY, costToNodeFromStart: 0, directionIndex: east, index: 0}}
	heap.Init(&openSet)
	gScoreMap := make(map[string]int)
	gScoreMap[fmt.Sprintf("%d,%d", startingX, startingY)] = 0
	fScoreMap := make(map[string]int)
	fScoreMap[fmt.Sprintf("%d,%d", startingX, startingY)] = 0

	for len(openSet) > 0 {
		currentNode := heap.Pop(&openSet).(*Node)
		if mapMatrix[currentNode.x][currentNode.y] == "E" {
			return currentNode.costToNodeFromStart
		}

		for _, directionV := range directionVectors {
			neighborX := currentNode.x + directionV[0]
			neighborY := currentNode.y + directionV[1]
			if neighborX < 0 || neighborX >= len(mapMatrix) || neighborY < 0 || neighborY >= len(mapMatrix[0]) {
				continue
			}
			if mapMatrix[neighborX][neighborY] == "#" {
				continue
			}

			tentativeGScore := gScoreMap[fmt.Sprintf("%d,%d", currentNode.x, currentNode.y)] + heuristicFunc(currentNode.directionIndex, neighborX-currentNode.x, neighborY-currentNode.y)
			_, ok := gScoreMap[fmt.Sprintf("%d,%d", neighborX, neighborY)]
			if !ok || tentativeGScore < gScoreMap[fmt.Sprintf("%d,%d", neighborX, neighborY)] {
				gScoreMap[fmt.Sprintf("%d,%d", neighborX, neighborY)] = tentativeGScore
			}

			heap.Push(&openSet, &Node{
				x:                   neighborX,
				y:                   neighborY,
				costToNodeFromStart: gScoreMap[fmt.Sprintf("%d,%d", neighborX, neighborY)],
				directionIndex:      currentNode.directionIndex,
				index:               len(openSet)})
		}
	}

	return -1
}

func traverse(mazeMatrix [][]string, x, y int, scoreMatrix [][]int, visitedMatrix [][]bool, currentCost int, directionIndex int) {
	if currentCost < scoreMatrix[x][y] {
		scoreMatrix[x][y] = currentCost
	}
	visitedMatrix[x][y] = true
	// otherwise, recursively traverse each direction if that direction is in range and not a #
	// question is how to calculate traversal cost? we need to know in each recursion what direction we are facing

	// continue north if possible
	if x > 0 && mazeMatrix[x-1][y] != "#" && !visitedMatrix[x-1][y] {
		northCost := currentCost + 1

		if directionVectors[directionIndex][0] != -1 {
			northCost += 1000
		}
		traverse(mazeMatrix, x-1, y, scoreMatrix, copyVisitedMatrix(visitedMatrix), northCost, north)
	}

	// continue east if possible
	if y < len(mazeMatrix[x])-1 && mazeMatrix[x][y+1] != "#" && !visitedMatrix[x][y+1] {
		eastCost := currentCost + 1
		if directionVectors[directionIndex][1] != 1 {
			eastCost += 1000
		}
		traverse(mazeMatrix, x, y+1, scoreMatrix, copyVisitedMatrix(visitedMatrix), eastCost, east)
	}

	// continue south if possible
	if x < len(mazeMatrix)-1 && mazeMatrix[x+1][y] != "#" && !visitedMatrix[x+1][y] {
		southCost := currentCost + 1
		if directionVectors[directionIndex][0] != 1 {
			southCost += 1000
		}
		traverse(mazeMatrix, x+1, y, scoreMatrix, copyVisitedMatrix(visitedMatrix), southCost, south)
	}

	// continue west if possible
	if y > 0 && mazeMatrix[x][y-1] != "#" && !visitedMatrix[x][y-1] {
		westCost := currentCost + 1
		if directionVectors[directionIndex][1] != -1 {
			westCost += 1000
		}
		traverse(mazeMatrix, x, y-1, scoreMatrix, copyVisitedMatrix(visitedMatrix), westCost, west)
	}

}

func copyVisitedMatrix(visitedMatrix [][]bool) [][]bool {
	newMatrix := make([][]bool, len(visitedMatrix))
	for i := range visitedMatrix {
		newMatrix[i] = make([]bool, len(visitedMatrix[i]))
		copy(newMatrix[i], visitedMatrix[i])
	}
	return newMatrix
}

func sanityCheck() {
	testMaze := `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

	testMazeMatrix := buildMazeMatrix(testMaze)
	// testScoreMatrix := buildScoreMatrix(testMazeMatrix)
	startingX, startingY := getStartingPoint(testMazeMatrix)
	// traverse(testMazeMatrix, startingX, startingY, testScoreMatrix, buildVisitedMatrix(testMazeMatrix), 0, east)

	// cost := findCheapestRoute(testMazeMatrix, testScoreMatrix)
	cost := aStarTraversal(testMazeMatrix, startingX, startingY)
	if cost != 11048 {
		panic("from test data, cost should be 11048 but is " + fmt.Sprint(cost))
	}
}

func printScoreMatrix(testScoreMatrix [][]int) {
	// Find the maximum width needed
	maxWidth := 0
	for _, row := range testScoreMatrix {
		for _, score := range row {
			if score == math.MaxInt {
				if 1 > maxWidth {
					maxWidth = 1
				}
			} else {
				width := len(fmt.Sprint(score))
				if width > maxWidth {
					maxWidth = width
				}
			}
		}
	}

	fmt.Println("Score Matrix:")
	for _, row := range testScoreMatrix {
		for _, score := range row {
			if score == math.MaxInt {
				fmt.Printf("%*s", maxWidth+2, "∞")
			} else {
				fmt.Printf("%*d", maxWidth+2, score)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func findCheapestRoute(mazeMatrix [][]string, scoreMatrix [][]int) int {
	for x, line := range mazeMatrix {
		for y, spot := range line {
			if spot == "E" {
				return scoreMatrix[x][y]
			}
		}
	}
	return -1
}

func buildMazeMatrix(maze string) [][]string {
	mazeMatrix := [][]string{}
	for _, line := range strings.Split(maze, "\n") {
		mazeMatrix = append(mazeMatrix, strings.Split(line, ""))
	}
	return mazeMatrix
}

func getStartingPoint(maze [][]string) (int, int) {
	for x, line := range maze {
		for y, spot := range line {
			if spot == "S" {
				return x, y
			}
		}
	}
	return -1, -1
}
