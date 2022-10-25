package sort3


func QuickSort(a []int) []int {
	if len(a) < 1 {
		return a
	}
	recursionQuickSort(a, 0, len(a)-1)
	return a
}


func recursionQuickSort(a []int, l, r int) {
	if l<0 || r>len(a)-1 || l>=r {
		return
	}

	pivot := a[l]
	left, right := l, r
	for left < right {
		for left < right && a[right] >= pivot {
			right --
		}

		for left < right && a[left] <= pivot {
			left ++
		}

		a[left], a[right] = a[right], a[left]
	}

	a[left], a[l] = a[l], a[left]

	recursionQuickSort(a, l, left - 1)
	recursionQuickSort(a, left + 1, r)
}
