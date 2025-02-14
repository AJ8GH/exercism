package minesweeper

import "fmt"

const (
	mine = '*'
)

var neighbours = []func(int, int) (int, int){
	func(i int, j int) (int, int) { return i - 1, j },
	func(i int, j int) (int, int) { return i + 1, j },
	func(i int, j int) (int, int) { return i, j + 1 },
	func(i int, j int) (int, int) { return i, j - 1 },
	func(i int, j int) (int, int) { return i - 1, j + 1 },
	func(i int, j int) (int, int) { return i - 1, j - 1 },
	func(i int, j int) (int, int) { return i + 1, j + 1 },
	func(i int, j int) (int, int) { return i + 1, j - 1 },
}

// Annotate returns an annotated board
func Annotate(board []string) (outBoard []string) {
	for i, row := range board {
		outRow := ""
		for j, square := range row {
			switch square {
			case mine:
				outRow += string(square)
			default:
				score := score(i, j, board)
				if score > 0 {
					outRow += fmt.Sprintf("%d", score)
				} else {
					outRow += " "
				}
			}
		}
		outBoard = append(outBoard, outRow)
	}
	return outBoard
}

func score(i, j int, board []string) (outScore int) {
	for _, v := range neighbours {
		newI, newJ := v(i, j)
		switch {
		case newI < 0 || newJ < 0 || newI >= len(board) || newJ >= len(board[newI]):
			continue
		case board[newI][newJ] == mine:
			outScore++
		}
	}
	return outScore
}
