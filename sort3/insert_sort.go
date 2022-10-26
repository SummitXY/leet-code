package sort3

func InsertSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	for i:=1;i<len(a);i++ {
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
