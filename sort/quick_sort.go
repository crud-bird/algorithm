package sort

import "fmt"

//QuickSort 快排
func QuickSort(a []int64, idxLow, idxHigh int) {
	if idxHigh-idxLow < 1 {
		return
	}

	low, high := idxLow, idxHigh

	for low < high {
		for low < high && a[low] <= a[idxLow] {
			low++
		}
		for low < high && a[high] >= a[idxLow] {
			high--
		}
		if low < high {
			a[low], a[high] = a[high], a[low]
		} else {
			if a[low] < a[idxLow] {
				a[low], a[idxLow] = a[idxLow], a[low]
			}
		}
	}
	if high != low {
		fmt.Printf("%d != %d", low, high)
	}
	QuickSort(a, idxLow, low-1)
	QuickSort(a, high, idxHigh)
}
