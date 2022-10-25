package sort
//
//func MergeSort(a []int) []int {
//
//	mergeSortSort(a, 0, len(a)-1)
//	return a
//}
//
//func mergeSortSort(a []int, l, r int) {
//	if l >= r {
//		return
//	}
//
//	mid := (l + r)/2 // 左边的最右边界
//	mergeSortSort(a, l, mid)
//	mergeSortSort(a, mid+1, r)
//	merge(a, l, mid, r)
//}
//
//func merge(a []int, l, mid, r int) {
//
//	la := make([]int, 0, mid-l+1)
//	ra := make([]int, 0, r-mid)
//	copy(la, )
//	la := a[l:mid+1]
//	ra := a[mid+1:r+1]
//
//	lidx:=0
//	ridx:=0
//
//	for i:=l;i<=r;i++ {
//		if lidx > len(la) - 1 {
//			// la 空了
//			a[i] = ra[ridx]
//			ridx ++
//		} else if ridx > len(ra) - 1 {
//			// ra 空了
//			a[i] = la[lidx]
//			lidx++
//		} else {
//			// la ra 都没空
//			if la[lidx] < ra[ridx] {
//				a[i] = la[lidx]
//				lidx ++
//			} else {
//				a[i] = ra[ridx]
//				ridx ++
//			}
//		}
//	}
//}
