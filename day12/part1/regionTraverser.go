package main

import (
	"strings"
)

type regionTraverser struct {
	farmMatrix    [][]string
	exploredPlots [][]int
	regions       []region
}

func newRegionTraverser(farm string) *regionTraverser {
	farmRows := strings.Split(farm, "\n")
	farmMatrix := make([][]string, len(farmRows))
	for i, row := range farmRows {
		farmMatrix[i] = strings.Split(row, "")
	}
	exploredPlots := make([][]int, len(farmMatrix))
	for i := 0; i < len(exploredPlots); i++ {
		exploredPlots[i] = make([]int, len(farmMatrix[i]))
	}

	return &regionTraverser{
		farmMatrix:    farmMatrix,
		exploredPlots: exploredPlots,
		regions:       make([]region, 0),
	}
}

func (rt *regionTraverser) calculateFenceCost() int {
	for row := 0; row < len(rt.farmMatrix); row++ {
		for col := 0; col < len(rt.farmMatrix[row]); col++ {
			if rt.exploredPlots[row][col] == 1 {
				continue
			}
			newRegion := newRegion(rt.farmMatrix[row][col])
			rt.createRegion(row, col, newRegion)
			rt.regions = append(rt.regions, *newRegion)
		}
	}

	totalCost := 0
	for _, region := range rt.regions {
		// increment total cost per cost of region
		totalCost += len(region.plots) * region.calculatePerimeterCost()
	}
	return totalCost
}

func (rt *regionTraverser) createRegion(row int, col int, region *region) {
	if rt.farmMatrix[row][col] != region.name || rt.exploredPlots[row][col] == 1 {
		return
	}
	region.addPlot(row, col)
	rt.exploredPlots[row][col] = 1

	// Check up
	if row > 0 {
		rt.createRegion(row-1, col, region)
	}
	// Check down
	if row < len(rt.farmMatrix)-1 {
		rt.createRegion(row+1, col, region)
	}
	// Check left
	if col > 0 {
		rt.createRegion(row, col-1, region)
	}
	// Check right
	if col < len(rt.farmMatrix[row])-1 {
		rt.createRegion(row, col+1, region)
	}

}
