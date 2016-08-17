//Greed dice game kata.
package greedgo

import "sort"

type Roll []int

func (roll Roll) startsWith(anotherRoll ...int) bool {
	if len(anotherRoll) > len(roll) {
		return false
	}

	for i, value := range anotherRoll {
		if roll[i] != value {
			return false
		}
	}

	return true
}

func GreedMatching(dice []int) int {
	sort.Ints(dice)
	score := 0
	for len(dice) > 0 {
		score, dice = compute(score, dice)
	}
	return score
}

func compute(score int, sortedDice Roll) (int, Roll) {
	switch {
	case sortedDice.startsWith(1, 1, 1):
		return score + 1000, sortedDice[3:]
	case sortedDice.startsWith(1):
		return score + 100, sortedDice[1:]
	case sortedDice.startsWith(2, 2, 2):
		return score + 100*2, sortedDice[3:]
	case sortedDice.startsWith(3, 3, 3):
		return score + 100*3, sortedDice[3:]
	case sortedDice.startsWith(4, 4, 4):
		return score + 100*4, sortedDice[3:]
	case sortedDice.startsWith(5, 5, 5):
		return score + 100*5, sortedDice[3:]
	case sortedDice.startsWith(6, 6, 6):
		return score + 100*6, sortedDice[3:]
	case sortedDice.startsWith(5):
		return score + 50, sortedDice[1:]
	}
	return score, sortedDice[1:]
}