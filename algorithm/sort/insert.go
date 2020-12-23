package sort

func insertAsc(data []int) {
	for i := 1; i < len(data); i++ {
		tmp := data[i]
		k := i
		for j := i; j > 0 && tmp < data[j-1]; j-- {
			data[j] = data[j-1]
			k = j - 1
		}
		data[k] = tmp
	}

}
