package sort

func SelectSort(a []int) []int {
	for i:=0;i<len(a);i++ {
		m := a[0]
		p := 0
		for j:=0;j<len(a)-i;j++ {
			if a[j] > m {
				m = a[j]
				p = j
			}
		}
		a[p], a[len(a)-1-i] = a[len(a)-1-i], a[p]
	}

	return a
}

