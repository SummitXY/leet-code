package sort

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	quickSort(a, 0, len(a)-1)
	return a
}

// {3,7,5,9,4,6,2}
// {3,2,5,9,4,6,7}
// {}
func quickSort(a []int, l, r int) {
	if l >= r {
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

	mid := left
	if l < mid {
		quickSort(a, l, mid-1)
	}
	if r > mid {
		quickSort(a, mid+1, r)
	}
}