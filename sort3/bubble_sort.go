package sort3

func BubbleSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	for i:=len(a)-1;i>0;i-- {
		for j:=1;j<=i;j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}

	return a
}
