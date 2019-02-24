package main

import (
	"fmt"
)

const (
	MAX_N = 50
)

func main() {
	var n, m int
	fmt.Print("Please Input the number, expected sum:")
	_, err := fmt.Scanf("%d,%d", &n, &m)
	if err != nil {
		fmt.Printf("Scanf n and n error %v", err)
	}
	fmt.Println("Input number:", n, "Expected sum:", m)
	in := make([]int, n, n)
	fmt.Print("Please Input the array:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d,%d", &in[i])
	}
	fmt.Println("Input arrays:", in)

	f, combs := HasExpectedSumupCombs(m, in)
	fmt.Printf("Comb Exist:%v, Combs:%v \n", f, combs)
}

/*HasExpectedSumupCombs: check the roll 4 time can add sumup to expected value*/
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
