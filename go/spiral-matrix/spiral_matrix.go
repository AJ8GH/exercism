package spiralmatrix

type loc struct {
	i, j int
}

const (
	l = "l"
	r = "r"
	u = "u"
	d = "d"
)

func SpiralMatrix(size int) [][]int {
	spiral := [][]int{}
	for i := 0; i < size; i++ {
		spiral = append(spiral, make([]int, size))
	}

	curLoc := loc{i: 0, j: 0}
	dir := r

	for i := 1; i <= size*size; i++ {
		spiral[curLoc.i][curLoc.j] = i
		curLoc, dir = updateLoc(curLoc, spiral, size, dir)
	}

	return spiral
}

func updateLoc(curLoc loc, spiral [][]int, size int, dir string) (loc, string) {
	newLoc := nextLoc(dir, curLoc)
	if !validLoc(newLoc, size, spiral) {
		newDir := nextDir(dir)
		return nextLoc(newDir, curLoc), newDir
	}
	return newLoc, dir
}

func validLoc(newLoc loc, size int, spiral [][]int) bool {
	return newLoc.i >= 0 &&
		newLoc.j >= 0 &&
		newLoc.i < size &&
		newLoc.j < size &&
		spiral[newLoc.i][newLoc.j] == 0
}

func nextDir(curDir string) string {
	switch curDir {
	case l:
		return u
	case r:
		return d
	case d:
		return l
	case u:
		return r
	default:
		return curDir
	}
}

func nextLoc(curDir string, curLoc loc) loc {
	switch curDir {
	case l:
		return loc{i: curLoc.i, j: curLoc.j - 1}
	case r:
		return loc{i: curLoc.i, j: curLoc.j + 1}
	case d:
		return loc{i: curLoc.i + 1, j: curLoc.j}
	case u:
		return loc{i: curLoc.i - 1, j: curLoc.j}
	default:
		return curLoc
	}
}

func get(spiral [][]int, l loc) int {
	return spiral[l.i][l.j]
}

/*

    1  2  3 4  5
   16 17 18 19 6
   15 24 25 20 7
   14 23 22 21 8
   13 12 11 10 9
*/
