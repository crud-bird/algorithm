package sort

//MaxHeapify 维护最大堆性质
func MaxHeapify(a []int64, top int, last int) {
	l := top * 2
	r := top*2 + 1
	if l > last {
		return
	} else if r > last {
		if a[top] >= a[l] {
			return
		}
		a[top], a[l] = a[l], a[top]
		MaxHeapify(a, l, last)
		return
	} else {
		if a[top] >= a[l] && a[top] >= a[r] {
			return
		} else if a[l] > a[top] && a[l] > a[r] {
			a[top], a[l] = a[l], a[top]
			MaxHeapify(a, l, last)
		} else {
			a[top], a[r] = a[r], a[top]
			MaxHeapify(a, r, last)
		}
	}
}

//BuildHeap 建堆
func BuildHeap(a []int64) {
	aLen := len(a)
	top := aLen / 2
	for top > 0 {
		MaxHeapify(a, top, aLen-1)
		top--
	}
}

//HeapSort 堆排序
func HeapSort(a []int64) {
	BuildHeap(a)
	for i := len(a) - 1; i > 1; i-- {
		a[1] = a[1] + a[i]
		a[i] = a[1] - a[i]
		a[1] = a[1] - a[i]
		MaxHeapify(a, 1, i-1)
	}
}
