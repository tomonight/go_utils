package sort

func selectAsc(data []int) {
	for i := 0; i < len(data)-1; i++ {
		sub := i
		min := data[i]
		for j := i + 1; j < len(data); j++ {
			if data[j] < min {
				min = data[j]
				sub = j
			}
		}
		x := data[i]
		data[i] = min
		data[sub] = x
	}
}
