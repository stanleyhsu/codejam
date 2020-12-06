package leetcode

func maxProfit_ii(prices []int) int {
	n := len(prices)
	var greedyProfitSum int
	for i := 0; i < n-1; i++ {
		if profit := prices[i+1] - prices[i]; profit > 0 {
			greedyProfitSum += profit
		}
	}
	return greedyProfitSum
}
