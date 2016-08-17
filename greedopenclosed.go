//Greed dice game kata.
package greedgo

type Occurrencies map[int]int

func (occurrencies Occurrencies) remove (times, number int) {
	if occurrencies[number] > times {
		occurrencies[number] = occurrencies[number] - times
	} else {
		occurrencies[number] = 0
	}
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

func nItemsRule(toFind, scoreToAdd, nrOfItems int) func (int, Occurrencies) (int, Occurrencies) {
	return func (score int, occurrencies Occurrencies) (int, Occurrencies) {
		for ; occurrencies[toFind] >= nrOfItems; {
			score = score + scoreToAdd
			occurrencies.remove(nrOfItems, toFind)
		}
		return score, occurrencies
	} 
}