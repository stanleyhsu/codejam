package main

import (
	"fmt"
	"sort"
)

func main() {
	edges := []int{2, 3, 4, 5, 10}
	sort.Sort(sort.IntSlice(edges))
	fmt.Printf("Edges[%d]:%v\n", len(edges), edges)

	maxlen, maxedges := MaxTriangleLenAndEdges(edges[:])
	fmt.Printf("Exist?%v with edges:%v\n", maxlen, maxedges)
}

func rollgame() {
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
