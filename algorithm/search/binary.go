package search

func BinarySearch(arr []int, val int) int {
	return binarySearch(arr, val, 0, len(arr)-1)
}

func binarySearch(arr []int, val, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if val == arr[mid] {
		return mid
	}
	if val > arr[mid] {
		return binarySearch(arr, val, mid+1, right)
	} else {
		return binarySearch(arr, val, left, mid-1)
	}
}

func noRecursion(arr []int, val int) int {

	tmpL := len(arr) - 1
	left, right := 0, tmpL
	for left <= right {
		tmp := (left + right) >> 1
		if arr[tmp] == val {
			return tmp
		} else if arr[tmp] > val {
			right = tmp - 1
		} else {
			left = tmp + 1
		}
	}
	return -1
}
