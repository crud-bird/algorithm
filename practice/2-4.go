package practice

//FindReverseOrderNums 查找s中的逆序对，i < j 时，s[i]>s[j]，利用归并排序
func FindReverseOrderNums(s []int64) int {
	lens := len(s)
	if lens == 1 {
		return 0
	}

	lenL := lens / 2
	lenR := lens - lenL
	l := make([]int64, lenL)
	r := make([]int64, lenR)
	copy(l, s[:lenL])
	copy(r, s[lenL:])

	//逆序对数=左边数量+右边数量+左右间的数量
	count := FindReverseOrderNums(l)
	count += FindReverseOrderNums(r)

	//归并排序begin
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
	} //归并排序end

	//计算排好序的左右子数组之间的逆序对
	for i, j = lenL, lenR; i > 0 && j > 0; {
		if l[i-1] > r[j-1] {
			count += j
			i--
		} else {
			j--
		}
	}
	return count
}
