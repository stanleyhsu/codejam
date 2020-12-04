package leetcode

func sampleStats(count []int) []float64 {
	var minflag bool
	var min, max, mean, median float64

	var valueSum, countSum, modeIdx, countMax int
	var accumCount []accumCountType

	for i := 0; i <= 255; i++ {
		if count[i] != 0 {
			if !minflag {
				min = float64(i)
				minflag = true
			}
			max = float64(i)

			valueSum += i * count[i]
			countSum += count[i]
			accum := accumCountType{i, countSum}
			accumCount = append(accumCount, accum)

			if count[i] > countMax {
				countMax = count[i]
				modeIdx = i
			}
		}
	}
	mean = float64(valueSum) / float64(countSum)
	median = getMedian(accumCount, countSum)

	return []float64{min, max, mean, median, float64(modeIdx)}
}

type accumCountType struct {
	k     int
	accum int
}

func getMedian(ac []accumCountType, x int) float64 {
	n := len(ac)
	if x%2 == 0 {
		medianKl, medianKf := x/2, x/2+1
		for i := 0; i < n; i++ {
			if ac[i].accum >= medianKl && ac[i].accum >= medianKf {
				return float64(ac[i].k)
			}
			if ac[i].accum >= medianKl && ac[i].accum < medianKf {
				return float64(ac[i].k+ac[i+1].k) / 2.0
			}
		}
	} else {
		medianK := (x + 1) / 2
		for i := 0; i < n; i++ {
			if ac[i].accum >= medianK {
				return float64(ac[i].k)
			}
		}
	}
	return float64(0)
}
