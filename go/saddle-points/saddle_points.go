package matrix

import (
	"math"
	"strconv"
	"strings"
)

type Matrix struct {
	grid [][]int
}

type Pair struct {
	a, b int
}

type TreeLoc struct {
	i, j, val int
}

func New(s string) (*Matrix, error) {
	grid := [][]int{}
	for _, line := range strings.Split(s, "\n") {
		gridLine := []int{}
		for _, n := range strings.Split(line, " ") {
			d, _ := strconv.Atoi(n)
			gridLine = append(gridLine, d)
		}
		grid = append(grid, gridLine)
	}
	return &Matrix{grid: grid}, nil
}

func (m *Matrix) Saddle() (pairs []Pair) {
	if m.isEmpty() {
		return pairs
	}
	maxes := findMaxes(m.grid)
	for _, mx := range maxes {
		if isMin(mx, m.grid) {
			pairs = append(pairs, Pair{a: mx.i + 1, b: mx.j + 1})
		}
	}
	return pairs
}

func (m *Matrix) isEmpty() bool {
	if len(m.grid) == 1 && len(m.grid[0]) == 1 {
		return true
	}
	for _, row := range m.grid {
		if len(row) > 0 {
			return false
		}
	}
	return true
}

func findMaxes(grid [][]int) (locs []TreeLoc) {
	for i, row := range grid {
		mx := math.MinInt
		iLocs := []TreeLoc{}
		for j, tree := range row {
			switch {
			case tree > mx:
				mx = tree
				iLocs = append([]TreeLoc{}, TreeLoc{i: i, j: j, val: tree})
			case tree == mx:
				iLocs = append(iLocs, TreeLoc{i: i, j: j, val: tree})
			}
		}
		locs = append(locs, iLocs...)
	}
	return locs
}

func isMin(l TreeLoc, grid [][]int) bool {
	for _, row := range grid {
		if row[l.j] < l.val {
			return false
		}
	}
	return true
}
