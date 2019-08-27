package practice

//SumX 从有序数组s中查找和为x的两个数
func SumX(s []int64, x int64) (int, int) {
	i := 0
	j := len(s) - 1
	for i < j {
		sum := s[i] + s[j]
		if sum == x {
			return i, j
		} else if sum > x {
			j--
		} else {
			i++
		}
	}
	return -1, -1
}
