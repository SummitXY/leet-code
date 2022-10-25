package sort2

func QuickSort(a []int) []int {

	quick(a, 0, len(a)-1)
	return a
}

func quick(a []int, left, right int) {
	if left >= right {
		return
	}

	mid := quickSort(a, left, right)
	quick(a, left, mid-1)
	quick(a, mid+1, right)
}

func quickSort(a []int, left, right int) int {
	pivot := a[left]
	l := left
	r := right

	for l < r {
		for l < r && a[r] > pivot {
			r--
		}
		for l < r && a[l] < pivot {
			l++
		}
		a[l], a[r] = a[r], a[l]
	}

	a[l] = pivot
	return l
}
