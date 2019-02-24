package main

/*
MaxTriangleLenAndEdges max edges
*/
func MaxTriangleLenAndEdges(edges []int) (maxlen int, maxedges [3]int) {
	n := len(edges)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				total := edges[i] + edges[j] + edges[k]
				if edges[i]+edges[j] > edges[k] {
					if total > maxlen {
						maxlen = total
						maxedges = [3]int{edges[i], edges[j], edges[k]}
					}
				}
			}
		}
	}
	return
}
