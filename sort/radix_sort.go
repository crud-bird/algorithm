package sort

//RadixSort 基数排序， 内部基于计数排序
func RadixSort(a []int64, d, lenth int) []int64 {
	cur := make(map[int64]int64)
	for _, v := range a {
		cur[v] = v
	}
	b := make([]int64, 0)
	b = append(b, a...)
	for k := 0; int(k) < lenth; k++ {
		c := make([]int64, d)
		for _, v := range a {
			c[cur[v]%10]++
		}
		for i := 1; i < d; i++ {
			c[i] = c[i] + c[i-1]
		}
		for i := len(a); i > 0; i-- {
			v := cur[a[i-1]]
			c[v%10] = c[v%10] - 1
			b[c[v%10]] = a[i-1]
		}
		for i, v := range b {
			a[i] = v
		}
		for key, val := range cur {
			cur[key] = val / 10
		}
	}

	return a
}
