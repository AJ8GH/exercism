package rectangles

const (
	corner = '+'
	hSide  = '-'
	vSide  = '|'
)

func Count(diagram []string) (count int) {
	for i, row := range diagram {
		for j, v := range row {
			if v == corner {
				count += findRectangles(i, j, diagram)
			}
		}
	}

	return count
}

func findRectangles(i, j1 int, diagram []string) (count int) {
	// Find next corner

loop:
	for j2 := j1 + 1; j2 < len(diagram[i]); j2++ {
		v := diagram[i][j2]
		switch {
		case v != corner && v != hSide:
			break loop
		case v == corner:
			count += findRectanglesVertically(i, j1, j2, diagram)
		}
	}

	return count
}

func findRectanglesVertically(i, j1, j2 int, diagram []string) (count int) {
	j1Corners := map[int]bool{}
loop:
	for i1 := i + 1; i1 < len(diagram); i1++ {
		v := diagram[i1][j1]
		switch {
		case v != corner && v != vSide:
			break loop
		case v == corner:
			j1Corners[i1] = true
		}
	}

	j2Corners := map[int]bool{}
loop1:
	for i1 := i + 1; i1 < len(diagram); i1++ {
		v := diagram[i1][j2]
		switch {
		case v != corner && v != vSide:
			break loop1
		case v == corner:
			j2Corners[i1] = true
		}
	}

	for k := range j1Corners {
		if j2Corners[k] {
			for j3 := j1; j3 < j2; j3++ {
				v := diagram[k][j3]
				if v != corner && v != hSide {
					break
				}
				if j3 == j2-1 {
					count++
				}
			}
		}
	}
	return count
}
