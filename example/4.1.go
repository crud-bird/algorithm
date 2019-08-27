package example

//MaxSumOfSubarray 查找子数组的最大和,分治
func MaxSumOfSubarray(a []int, low, high int) (idxL, idxH, maxSum int) {
	if len(a) == 1 {
		idxH = high
		idxL = low
		maxSum = a[low]
		return
	}

	mid := (low + high) / 2

	//左边最大值
	LIdxL, LIdxH, LMaxSum := MaxSumOfSubarray(a, low, mid)
	//右边最大值
	RIdxL, RIdxH, RMaxSum := MaxSumOfSubarray(a, mid+1, high)
	//跨越两个子数组的情况
	CIdxL, CIdxH, CMaxSum := MaxSumOfCrossingSubarray(a, low, mid, high)

	if LMaxSum >= RMaxSum && LMaxSum >= CMaxSum {
		idxL = LIdxL
		idxH = LIdxH
		maxSum = LMaxSum
	} else if RMaxSum >= LMaxSum && RMaxSum >= CMaxSum {
		idxH = RIdxH
		idxL = RIdxL
		maxSum = RMaxSum
	} else {
		idxH = CIdxH
		idxL = CIdxL
		maxSum = CMaxSum
	}
	return
}

//MaxSumOfCrossingSubarray 跨越两个子数组的情况
func MaxSumOfCrossingSubarray(a []int, low, mid, high int) (idxL, idxH, maxSum int) {
	maxL, maxR := -100000, -100000
	sumL, sumR := 0, 0

	//左边最大和
	for i := mid; i >= low; i-- {
		sumL += a[i]
		if sumL > maxL {
			maxL = sumL
			idxL = i
		}
	}

	//右边最大和
	for i := mid; i <= high; i++ {
		sumR += a[i]
		if sumR > maxR {
			maxR = sumR
			idxH = i
		}
	}

	maxSum = maxL + maxR
	return
}
