package sort3

func SelectSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	for i:=len(a)-1;i>0;i-- {
		max := 0
		for j:=0;j<=i;j++ {
			if a[j] > a[max] {
				max = j
			}
		}
		a[i], a[max] = a[max], a[i]
	}

	return a
}
