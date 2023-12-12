package card

import "fmt"

var tokenStrengths = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type Card struct {
	Token    string
	Strength int
}

func NewCard(token string, usesJokers bool) Card {
	return Card{
		Token:    token,
		Strength: getTokenStrength(token, usesJokers),
	}
}

func getTokenStrength(token string, usesJokers bool) int {
	if usesJokers && token == "J" {
		return 1
	}

	if strength, ok := tokenStrengths[token]; ok {
		return strength
	}
	panic(fmt.Sprintf("Unknown token: %v", token))
}
