package queenattack

import (
	"errors"
	"math"
)

const (
	minRank = '1'
	maxRank = '8'
	minCol  = 'a'
	maxCol  = 'h'
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition || len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errors.New("invalid")
	}

	whiteColumn, whiteRank := rune(whitePosition[0]), rune(whitePosition[1])
	blackColumn, blackRank := rune(blackPosition[0]), rune(blackPosition[1])

	if whiteRank < minRank ||
		whiteRank > maxRank ||
		blackRank < minRank ||
		blackRank > maxRank ||
		whiteColumn < minCol ||
		whiteColumn > maxCol ||
		blackColumn < minCol ||
		blackColumn > maxCol {
		return false, errors.New("invalid")
	}

	if whiteRank == blackRank || whiteColumn == blackColumn {
		return true, nil
	}

	rankComp := math.Abs(float64(whiteRank - blackRank))
	colComp := math.Abs(float64(whiteColumn - blackColumn))

	return rankComp == colComp, nil
}
