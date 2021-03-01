package sort

func shellAsc(data []int) {
	group := len(data) / 2
	for group >= 1 {
		for i := 0; i < len(data)-group; i += group {
			tmp := data[i+group]
			for j := i; j >= 0; j -= group {
				if tmp < data[j] {

					data[j+group] = data[j]
					data[j] = tmp
				}
			}
		}
		group = group / 2
	}
}
