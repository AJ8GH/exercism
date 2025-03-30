package stateoftictactoe

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
	x             = "X"
	o             = "O"
	empty         = " "
	xWin          = "XXX"
	oWin          = "OOO"
)

func StateOfTicTacToe(board []string) (State, error) {
	cols := []string{"", "", ""}
	diags := []string{"", ""}
	ongoing := false
	numX := 0
	numO := 0

	for i, row := range board {
		for j, cell := range row {
			switch string(cell) {
			case empty:
				ongoing = true
			case x:
				numX++
			case o:
				numO++
			}
			cols[j] += string(cell)
			handleDiags(i, j, diags, string(cell))
		}
	}

	xWins := numWins(diags, isXWin) + numWins(board, isXWin) + numWins(cols, isXWin)
	oWins := numWins(diags, isOWin) + numWins(board, isOWin) + numWins(cols, isOWin)
	switch {
	case numO-numX > 0 || numX-numO > 1 || xWins > 0 && oWins > 0:
		return "", errors.New("bad board")
	case xWins > 0 || oWins > 0:
		return Win, nil
	case ongoing:
		return Ongoing, nil
	default:
		return Draw, nil
	}
}

func numWins(lines []string, isWin func(string) bool) (wins int) {
	for _, line := range lines {
		if isWin(line) {
			wins++
		}
	}
	return wins
}

func isOWin(line string) bool {
	if line == oWin {
		return true
	}
	return false
}

func isXWin(line string) bool {
	if line == xWin {
		return true
	}
	return false
}

func handleDiags(i, j int, diags []string, cell string) []string {
	switch {
	case i == 1 && j == 1:
		diags[0] += cell
		diags[1] += cell
	case i == j:
		diags[0] += cell
	case (i == 0 && j == 2) || (i == 2 && j == 0):
		diags[1] += cell
	}

	return diags
}
