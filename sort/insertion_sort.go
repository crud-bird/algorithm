package sort

//InsertionSort 插入排序
func InsertionSort(s []int64) []int64 {
	len := len(s)
	for i := 1; i < len; i++ {
		cur := s[i]
		j := i - 1
		for j >= 0 && s[j] > cur {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = cur
	}
	return s
}
