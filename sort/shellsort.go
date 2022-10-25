package sort

func ShellSort(a []int) []int {

	for gap:=len(a)/2;gap>=1;gap/=2 {
		for i:=gap;i<len(a);i++ {
			for j:=i;j>0&&j-gap>=0;j-=gap {
				if a[j] >= a[j-gap] {
					break
				} else {
					a[j], a[j-gap] = a[j-gap], a[j]
				}
			}
		}
	}

	return a
}
