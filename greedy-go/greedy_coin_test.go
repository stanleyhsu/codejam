package greedy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinCoinSeries(t *testing.T) {
	golden := []struct {
		name          string
		target        int
		valuesEnum    []int
		wantCoinCount int
		wantCoins     []int
	}{
		{"7-5", 7, []int{5}, -1, nil},
		{"12-521", 12, []int{5, 2, 1}, 3, []int{5, 5, 2}},
		//  {"11-53", 11, []int{5, 3}, 3, []int{5, 3, 3}}, //FIXME
	}
	// Iterate over test cases.
	for _, tt := range golden {
		t.Run(tt.name, func(t *testing.T) {
			counter, coins := getMinCoinSeries(tt.target, tt.valuesEnum)
			assert.Equal(t, tt.wantCoinCount, counter, "coin counter should be equal")
			assert.Equal(t, tt.wantCoins, coins, "coin array should be equal")
		})
	}
}

func TestGetMinCoinSeriesLoop(t *testing.T) {
	golden := []struct {
		name          string
		target        int
		valuesEnum    []int
		wantCoinCount int
		wantCoins     []int
	}{
		{"7-5", 7, []int{5}, -1, nil},
		{"12-521", 12, []int{5, 2, 1}, 3, []int{5, 5, 2}},
		{"13-521", 13, []int{5, 2, 1}, 4, []int{5, 5, 2, 1}},

		{"11-53", 11, []int{5, 3}, 3, []int{5, 3, 3}},
		{"14-53", 14, []int{5, 3}, 4, []int{5, 3, 3, 3}},
		{"14-10-7-1", 14, []int{10, 7, 1}, 2, []int{7, 7}},
	}
	// Iterate over test cases.
	for _, tt := range golden {
		t.Run(tt.name, func(t *testing.T) {
			counter := getMinCoinCountOfValue(tt.target, tt.valuesEnum)
			assert.Equal(t, tt.wantCoinCount, counter, "coin counter should be equal")
			//assert.Equal(t, tt.wantCoins, coins, "coin array should be equal")
		})
	}
}
