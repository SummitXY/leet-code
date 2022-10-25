package sort2

func InsertSort(a []int) []int {
	// time : n^2
	// spave : 1
	if len(a) < 2 {
		return a
	}

	for i:=0;i<len(a);i++ {
		for j:=i;j>0;j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}

	return a
}
