package sort

func shellAsc(data []int) {
	for step := len(data) / 2; step > 0; step = step / 2 {
		for i := step; i < len(data); i++ {
			for j := i - step; j >= 0 && data[j+step] < data[j]; j -= step {
				data[j+step], data[j] = data[j], data[j+step]
			}
		}
	}
}
