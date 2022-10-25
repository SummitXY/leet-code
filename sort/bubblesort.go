package sort


func BubbleSort(a []int) []int {

	for i:=0;i<len(a);i++ {
		for j:=0;j<len(a)-1-i;j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}