package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64((score - 10)) / float64(2)))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return roll() + roll() + roll()
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Ability(),
		Ability(),
		Ability(),
		Ability(),
		Ability(),
		Ability(),
		10,
	}
	c.Hitpoints += Modifier(c.Constitution)
	return c
}

func roll() int {
	return rand.Intn(6) + 1
}
