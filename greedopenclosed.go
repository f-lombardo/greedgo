//Greed dice game kata.
package greedgo

import "math"

type Occurrencies map[int]int

func (occurrencies Occurrencies) remove (times, number int) {
	occurrencies[number] = int(math.Max(float64(occurrencies[number] - times), float64(0)))
}


func GreedOpenClosed(dice []int) int {
	var rules [] func (int, Occurrencies) (int, Occurrencies)
	rules = append(rules, nItemsRule(1, 1000, 3))
	rules = append(rules, nItemsRule(1, 100, 1))	
	rules = append(rules, nItemsRule(2, 200, 3))
	rules = append(rules, nItemsRule(3, 300, 3))		
	rules = append(rules, nItemsRule(4, 400, 3))	
	rules = append(rules, nItemsRule(5, 500, 3))
	rules = append(rules, nItemsRule(5,  50, 1))	
	rules = append(rules, nItemsRule(6, 600, 3))	
	occurrencies := computeOccurrencies(dice)
	score := 0
	for _, rule := range rules {
		score, occurrencies = rule(score, occurrencies)
	}
	return score
}

func computeOccurrencies(dice []int) Occurrencies {
	occurrencies := Occurrencies(make(map[int]int))
	for _, value := range dice {
		occurrencies[value] = occurrencies[value] + 1
	}
	return occurrencies
}

func nItemsRule(toFind, score, nrOfItems int) func (int, Occurrencies) (int, Occurrencies) {
	return func (previousScore int, occurrencies Occurrencies) (int, Occurrencies) {
		for ; occurrencies[toFind] >= nrOfItems; {
			previousScore = previousScore + score
			occurrencies.remove(nrOfItems, toFind)
		}
		return previousScore, occurrencies
	} 
}