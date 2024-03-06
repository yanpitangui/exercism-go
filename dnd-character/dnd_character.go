package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
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
	return int(math.Floor((float64(score) - 10) / 2))

}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rollDice := func() int { return rand.Intn(6) + 1 }
	rolled := []int{}
	for i := 0; i < 6; i++ {
		rolled = append(rolled, rollDice())
	}
	sum := 0
	slices.Sort(rolled)

	for i := len(rolled) - 1; i >= 3; i-- {
		sum += rolled[i]
	}

	return sum

}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	con := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: con,
		Hitpoints:    10 + Modifier(con),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
}
