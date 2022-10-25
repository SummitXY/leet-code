package sort

func QuickSortTopK(a []int, k int) int {

	quickSortReturn(a, 0, len(a)-1, k)
	return a[len(a)-k]
}

func quickSortReturn(a []int, left, right, k int) {
	if left >= right {
		return
	}

	pivot := a[left]
	l, r := left, right
	for l < r {
		for l < r && a[r] >= pivot {
			r--
		}
		for l < r && a[l] <= pivot {
			l++
		}
		a[l], a[r] = a[r], a[l]
	}
	a[l], a[left] = a[left], a[l]

	mid := l
	if len(a) - mid == k {
		return
	} else if len(a) - mid > k {
		// k在右半部
		if mid < right {
			quickSortReturn(a, mid+1, right, k)
		}
	} else {
		// k在左半部
		if left < mid {
			quickSortReturn(a, left, mid-1, k)
		}
	}
}
