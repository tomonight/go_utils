package sort

func QuickAsc(arr []int) {
	quickAsc(arr, 0, len(arr)-1)
}

//
func quickAsc(arr []int, left, right int) {
	if left == right || right == 0 {
		return
	}
	l, r := left, right
	if l < r {
		k := partition(arr, l, r)
		quickAsc(arr, left, k-1)
		quickAsc(arr, k+1, right)
	}
	//po := left
	//tmp := arr[left]
	// l = l+1
	//for l < r {
	//	for ;l < r && arr[l] < tmp;l++{
	//	}
	//	for ;l < r && arr[r] > tmp;r--{
	//	}
	//	arr[l],arr[r] = arr[r],arr[l]
	//}
	//arr[left],arr[r - 1] = arr[r - 1],arr[left]
	//if left < r - 1{
	//	quickAsc(arr,left,r-1)
	//}
	//if right > r+1{
	//	quickAsc(arr,r+1,right)
	//}

}

func p(arr []int, left, right int) int {
	index := left + 1
	tmp := arr[left]

	for i := index + 1; i < right-left; i++ {
		if arr[i] < tmp {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[left], arr[index-1] = arr[index-1], arr[left]
	return index - 1
}

func partition(arr []int, left, right int) int {
	index := left + 1
	for i := left; i <= right; i++ {
		if arr[i] < arr[left] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index-1], arr[left] = arr[left], arr[index-1]
	return index - 1
}
