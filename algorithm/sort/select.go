package sort

func selectAsc(data []int) {
	for i := 0; i < len(data); i++ {
		min := data[i]
		t := i
		for j := i; j < len(data); j++ {
			if min > data[j] {
				min = data[j]
				t = j
			}
		}
		k := data[i]
		data[i] = min
		data[t] = k
	}
}
