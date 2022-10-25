package sort3


func QuickSort(a []int) []int {
	if len(a) < 1 {
		return a
	}
	recursionQuickSort(a, 0, len(a)-1)
	return a
}


func recursionQuickSort(a []int, l, r int) {
	pivot := a[l]
	left, right := l, r
	for left < right {
		for a[right] >= pivot {
			right --
		}

		for a[left] <= pivot {
			left ++
		}

		a[left], a[right] = a[right], a[left]
	}


}
