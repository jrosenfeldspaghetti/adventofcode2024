package main

import "fmt"

type region struct {
	plots map[string]int
	name  string
}

func newRegion(name string) *region {
	return &region{
		name:  name,
		plots: make(map[string]int),
	}
}

func (r *region) addPlot(x int, y int) {
	r.plots[fmt.Sprintf("%d,%d", x, y)] = 1
}

func (r *region) calculateSideCost() int {
	cornersFound := 0
	for coordinates := range r.plots {
		var x, y int
		fmt.Sscanf(coordinates, "%d,%d", &x, &y)
		cornersFoundForPlot := 0

		plotAboveInRegion := r.plots[fmt.Sprintf("%d,%d", x, y-1)]
		plotToRightInRegion := r.plots[fmt.Sprintf("%d,%d", x+1, y)]
		plotBelowInRegion := r.plots[fmt.Sprintf("%d,%d", x, y+1)]
		plotToLeftInRegion := r.plots[fmt.Sprintf("%d,%d", x-1, y)]
		// Check if top left corner is in region
		if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x-1, y-1)]; !plotInRegion {
			// not a corner iff one of up or to the right is in the region
			if plotAboveInRegion+plotToLeftInRegion != 1 {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}
		// Check if top right corner is in region
		if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x+1, y-1)]; !plotInRegion {
			if plotAboveInRegion+plotToRightInRegion != 1 {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}
		// Check if bottom right corner is in region
		if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x+1, y+1)]; !plotInRegion {
			if plotBelowInRegion+plotToRightInRegion != 1 {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}
		// Check if bottom left corner is in region
		if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x-1, y+1)]; !plotInRegion {
			if plotBelowInRegion+plotToLeftInRegion != 1 {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}

		// the inverse of the above can also be true and lead to a corner.
		// If neighbors are not in plot but you do have someone in your region to your corner, you still need a corner to be counted
		if plotAboveInRegion+plotToLeftInRegion == 0 {
			if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x-1, y-1)]; plotInRegion {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}
		if plotAboveInRegion+plotToRightInRegion == 0 {
			if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x+1, y-1)]; plotInRegion {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}
		if plotBelowInRegion+plotToRightInRegion == 0 {
			if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x+1, y+1)]; plotInRegion {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}
		if plotBelowInRegion+plotToLeftInRegion == 0 {
			if _, plotInRegion := r.plots[fmt.Sprintf("%d,%d", x-1, y+1)]; plotInRegion {
				cornersFound += 1
				cornersFoundForPlot++
			}
		}
	}

	return cornersFound
}
