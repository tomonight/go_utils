package sort

func heapAsc(arr []int) {
	tmpl := len(arr) - 1
	for j := tmpl - 1; j >= 0; j-- {
		for i := tmpl/2 - 1; i >= 0; i-- {
			adjust(arr, i, tmpl)
		}
		arr[0], arr[tmpl] = arr[tmpl], arr[0]
		tmpl--
	}
}

func adjust(arr []int, i, length int) {

	tmp := arr[i]
	for k := i*2 + 1; k < length; k = 2*k + 1 {
		if arr[k] < arr[k+1] {
			k++
		}
		if arr[k] > tmp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = tmp
}

func maxall(data ...int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	return max
}
