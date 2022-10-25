package sort2

func MergeSort(a []int) []int {
	mergeSort(a, 0, len(a)-1)
	return a
}

func mergeSort(a []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r)/2
	mergeSort(a, l, mid)
	mergeSort(a, mid+1, r)
	merge(a, l, r, mid)
}

func merge(a []int, l, r, mid int) {
	left := make([]int, mid-l+1)
	right := make([]int, r-mid)
	copy(left, a[l:mid+1])
	copy(right, a[mid+1:r+1])

	p := 0
	q := 0
	for i:=l;i<=r;i++ {
		if p > len(left)-1 {
			// left 没数了
			a[i] = right[q]
			q++
		} else if q > len(right)-1 {
			// right 没数了
			a[i] = left[p]
			p++
		} else {
			if left[p] < right[q] {
				a[i] = left[p]
				p++
			} else {
				a[i] = right[q]
				q++
			}
		}
	}
}

