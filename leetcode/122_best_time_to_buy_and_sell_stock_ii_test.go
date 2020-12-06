package leetcode

import "testing"

func Test_maxProfit_ii(t *testing.T) {
	golden := []struct {
		name   string
		prices []int
		want   int
	}{
		{"example1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"example2", []int{1, 2, 3, 4, 5}, 4},
		{"example3", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range golden {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_ii(tt.prices); got != tt.want {
				t.Errorf("maxProfit(%v) = %v, want %v", tt.prices, got, tt.want)
			}
		})
	}
}
