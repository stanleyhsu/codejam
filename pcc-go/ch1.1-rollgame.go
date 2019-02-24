package main

/*
HasExpectedSumupCombs check the roll 4 time can add sumup to expected value
*/
func HasExpectedSumupCombs(sum int, in []int) (f bool, combs [][4]int) {
	tmp := make([][4]int, 0, 10)
	for _, va := range in {
		for _, vb := range in {
			for _, vc := range in {
				for _, vd := range in {
					if va+vb+vc+vd == sum {
						f = true
						tmp = append(tmp, [4]int{va, vb, vc, vd})
					}
				}
			}
		}
	}
	return f, tmp
}
