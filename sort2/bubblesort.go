package sort2

func BubbleSort(a []int) []int {
	// time n^2
	// space 1

	if len(a) < 2 {
		return a
	}

	for i:=0;i<len(a);i++ {
		for j:=0;j<len(a)-i-1;j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
