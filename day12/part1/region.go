package main

import "fmt"

type region struct {
	plots []plot
	name  string
}
type plot struct {
	x int
	y int
}

func newRegion(name string) *region {
	return &region{
		name:  name,
		plots: make([]plot, 0),
	}
}

func (r *region) addPlot(x int, y int) {
	r.plots = append(r.plots, plot{x, y})
}

func (r *region) calculatePerimeterCost() int {
	// A node with 0 neighbors is of cost 4
	// A node with 1 neighbor is of cost 3
	// A node with 2 neighbors is of cost 2
	// A node with 3 neighbors is of cost 1
	// A node with 4 neighbors is of cost 0

	// For each plot, add to the neighborMap increment all cardinal neighbor values
	// then, for each plot, get the neighbor value from the map (2N)
	// This could have been done more efficiently by just having the plots in a map already, but hey you live and you learn

	neighborMap := make(map[string]int)
	totalCost := 0

	for _, plot := range r.plots {
		neighborMap[fmt.Sprintf("%d,%d", plot.x-1, plot.y)] += 1
		neighborMap[fmt.Sprintf("%d,%d", plot.x+1, plot.y)] += 1
		neighborMap[fmt.Sprintf("%d,%d", plot.x, plot.y-1)] += 1
		neighborMap[fmt.Sprintf("%d,%d", plot.x, plot.y+1)] += 1
	}

	for _, plot := range r.plots {
		totalCost += 4 - neighborMap[fmt.Sprintf("%d,%d", plot.x, plot.y)]
	}

	return totalCost
}
