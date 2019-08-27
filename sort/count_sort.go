package sort

//CountSort 计数排序,输入只能是一定范围内的整数
func CountSort(a []int64, k int) []int64 {
	c := make([]int64, k)
	b := make([]int64, len(a))

	for _, v := range a {
		c[v] = c[v] + 1
	}

	for i := 1; i < k; i++ {
		c[i] = c[i] + c[i-1]
	}

	for i := len(a); i > 0; i-- {
		v := a[i-1]
		b[c[v]-1] = v
		c[v] = c[v] - 1
	}

	return b
}
