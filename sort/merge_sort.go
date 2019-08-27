package sort

//MergeSort 归并排序
func MergeSort(s []int64) {
	lens := len(s)
	if lens == 1 {
		return
	}

	lenL := lens / 2
	lenR := lens - lenL
	l := make([]int64, lenL)
	r := make([]int64, lenR)
	copy(l, s[:lenL])
	copy(r, s[lenL:])
	MergeSort(l)
	MergeSort(r)

	i, j := 0, 0
	for i < lenL && j < lenR {
		if l[i] > r[j] {
			s[i+j] = r[j]
			j++
		} else {
			s[i+j] = l[i]
			i++
		}
	}

	if i == lenL {
		copy(s[i+j:], r[j:])
	} else {
		copy(s[i+j:], l[i:])
	}
	return
}
