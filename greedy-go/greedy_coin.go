package greedy

import "math"

func getMinCoinSeries(target int, coinValueEum []int) (counter int, coins []int) {
	length := len(coinValueEum)
	remain := target
	counter = 0
	minCoinsCounter := remain/coinValueEum[0] + 1
	maxCoinsCounter := remain/coinValueEum[length-1] + 1
	coins = make([]int, minCoinsCounter, maxCoinsCounter)

	for i := 0; i < length; i++ {
		if remain >= coinValueEum[i] {
			i_counter := remain / coinValueEum[i]
			for j := 0; j < i_counter; j++ {
				coins[j+counter] = coinValueEum[i]
			}
			counter += i_counter
			remain -= i_counter * coinValueEum[i]
			if remain == 0 {
				return counter, coins
			}
		}
	}
	return -1, nil

}

func getMinCoinCountOfValue(target int, coinValueEum []int) (counter int) {
	valueLength := len(coinValueEum)
	if 0 == valueLength {
		return -1
	}
	minCount := math.MaxInt8
	currentValue := coinValueEum[0]
	maxCount := target / currentValue

	for aCount := maxCount; aCount >= 0; aCount-- {
		rest := target - currentValue*aCount
		if rest == 0 {
			minCount = min(minCount, aCount)
			break
		}
		restCount := getMinCoinCountOfValue(rest, coinValueEum[1:])

		if restCount == -1 {
			if aCount == 0 {
				break
			}
			continue
		}
		minCount = min(minCount, aCount+restCount)
	}
	if minCount == math.MaxInt8 {
		return -1
	}
	return minCount
}
