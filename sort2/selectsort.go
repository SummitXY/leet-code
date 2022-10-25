package sort2

func SelectSort(a []int) []int {

	for i:=0;i<len(a);i++ {
		max := a[0]
		p := 0
		for j:=0;j<len(a)-i;j++ {
			if a[j] > max {
				max = a[j]
				p = j
			}
		}
		a[len(a)-1-i], a[p] = a[p], a[len(a)-1-i]
	}

	return a
}
