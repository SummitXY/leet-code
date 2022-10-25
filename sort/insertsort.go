package sort

func InsertSort(a []int) []int {

	for i:=0;i<len(a)-1;i++ {
		for j:=i+1;j>0;j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
	return a
}
