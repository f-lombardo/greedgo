package greedgo

import (	
	"testing"
)


func TestGreed(t *testing.T) {
	cases := []struct {expected int; input []int} {	
		{1000, []int{1, 1, 1, 6, 6,}},
		{1200, []int{1, 1, 1, 1, 1,}},
		{1100, []int{1, 1, 1, 5, 5,}},	
		{ 200, []int{2, 4, 2, 2, 3,}},
		{ 250, []int{2, 5, 2, 2, 3,}},			
		{ 550, []int{5, 5, 5, 5, 3,}},	
		{ 300, []int{1, 5, 5, 1, 6,}},	
		{ 400, []int{4, 4, 6, 4, 3,}},
		{ 600, []int{6, 4, 6, 6, 3,}},		
	}

	for _, aCase := range cases {
		actual := Greed(aCase.input)
		if actual != aCase.expected {
			t.Errorf("Greed(%d) == %d, expected %d", aCase.input, actual, aCase.expected)
		}
	}
}