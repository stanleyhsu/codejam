package leetcode

import (
	"math"
)

//func maxProfit(prices []int) int {
//	n := len(prices)
//
//	var greedyMaxProfit, onePassMaxProfit int
//	for i := 0; i < n-1; i++ {
//		for j := i + 1; j < n; j++ {
//			if profit := prices[j] - prices[i]; profit > onePassMaxProfit {
//				onePassMaxProfit = profit
//			}
//		}
//		if onePassMaxProfit > greedyMaxProfit {
//			greedyMaxProfit = onePassMaxProfit
//		}
//	}
//	return greedyMaxProfit
//}

func maxProfit(prices []int) int {
	n := len(prices)
	minPrice := math.MaxInt64
	maxProfit := 0
	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if profit := prices[i] - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
