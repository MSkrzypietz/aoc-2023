package hand

import (
	"github.com/MSkrzypietz/aoc-2023/day07/card"
)

type Hand struct {
	Cards    []card.Card
	Strength int
	Bid      int
}

func NewHand(input string, bid int, usesJokers bool) Hand {
	var cards []card.Card
	for _, ch := range input {
		cards = append(cards, card.NewCard(string(ch), usesJokers))
	}

	hand := Hand{Cards: cards, Bid: bid}

	if usesJokers {
		hand.Strength = getJokerHandStrength(cards)
	} else {
		hand.Strength = getHandStrength(cards)
	}

	return hand
}

func getJokerHandStrength(cards []card.Card) int {
	var jokerIndexes []int
	var distinctCards []card.Card
	for i, currCard := range cards {
		if currCard.Token == "J" {
			jokerIndexes = append(jokerIndexes, i)
		} else {
			distinctCards = append(distinctCards, currCard)
		}
	}

	if len(jokerIndexes) == 0 || len(jokerIndexes) == 5 {
		return getHandStrength(cards)
	}

	maxHandStrength := 0
	for _, cardCombination := range getJokerReplacementCombinations(distinctCards, len(jokerIndexes)) {
		testHand := append([]card.Card{}, distinctCards...)
		testHand = append(testHand, cardCombination...)
		handStrength := getHandStrength(testHand)
		if handStrength > maxHandStrength {
			maxHandStrength = handStrength
		}
	}

	return maxHandStrength
}

func getJokerReplacementCombinations(cards []card.Card, length int) [][]card.Card {
	var result [][]card.Card

	var generate func(start int, currentCombination []card.Card)
	generate = func(start int, currentCombination []card.Card) {
		if len(currentCombination) == length {
			temp := make([]card.Card, len(currentCombination))
			copy(temp, currentCombination)
			result = append(result, temp)
			return
		}

		for i := start; i < len(cards); i++ {
			currentCombination = append(currentCombination, cards[i])
			generate(i, currentCombination)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	generate(0, []card.Card{})
	return result
}

func getHandStrength(cards []card.Card) int {
	threeOfAKind := false
	pairs := 0
	for _, count := range getTokenCount(cards) {
		switch count {
		case 5:
			return 6
		case 4:
			return 5
		case 3:
			threeOfAKind = true
		case 2:
			pairs++
		}
	}

	if threeOfAKind && pairs > 0 {
		return 4
	} else if threeOfAKind {
		return 3
	} else if pairs > 1 {
		return 2
	} else if pairs > 0 {
		return 1
	}
	return 0
}

func getTokenCount(cards []card.Card) map[string]int {
	result := make(map[string]int)
	for _, currCard := range cards {
		if count, ok := result[currCard.Token]; ok {
			result[currCard.Token] = count + 1
		} else {
			result[currCard.Token] = 1
		}
	}
	return result
}

func (h Hand) Compare(other Hand) bool {
	if h.Strength != other.Strength {
		return h.Strength < other.Strength
	}

	for i, currCard := range h.Cards {
		if currCard.Strength != other.Cards[i].Strength {
			return currCard.Strength < other.Cards[i].Strength
		}
	}
	return false
}
