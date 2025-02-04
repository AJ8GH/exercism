package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix interface {
	Rows() [][]int
	Cols() [][]int
	Set(int, int, int) bool
}

type MatrixImpl struct {
	rows [][]int
}

func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	rows := [][]int{}
	if len(lines) == 0 {
		return nil, errors.New("empty")
	}
	width := -1
	for _, v := range lines {
		nums := strings.Split(v, " ")
		nums = filterEmpties(nums)
		if width != -1 && len(nums) != width {
			return nil, errors.New("uneven")
		}
		width = len(nums)
		row := []int{}
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		rows = append(rows, row)
	}
	return MatrixImpl{rows: rows}, nil
}

func (m MatrixImpl) Cols() [][]int {
	cols := [][]int{}
	h := len(m.rows)
	w := len(m.rows[h-1])

	for i := 0; i < w; i++ {
		col := []int{}
		for j := 0; j < h; j++ {
			col = append(col, m.rows[j][i])
		}
		cols = append(cols, col)
	}
	return cols
}

func (m MatrixImpl) Rows() [][]int {
	rowsCopy := make([][]int, len(m.rows))
	for i, v := range m.rows {
		rowCopy := make([]int, len(v))
		copy(rowCopy, v)
		rowsCopy[i] = rowCopy
	}
	return rowsCopy
}

func (m MatrixImpl) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= m.height() || col >= m.width() {
		return false
	}
	m.rows[row][col] = val
	return true
}

func filterEmpties(nums []string) []string {
	newNums := []string{}
	for _, v := range nums {
		if v != "" && v != " " {
			newNums = append(newNums, v)
		}
	}
	return newNums
}

func (m MatrixImpl) height() int {
	return len(m.rows)
}

func (m MatrixImpl) width() int {
	if m.height() == 0 {
		return 0
	}
	return len(m.rows[0])
}
