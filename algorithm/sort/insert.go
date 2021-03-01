package sort

//
func insertAsc(data []int) {
	for i := 0; i < len(data)-1; i++ {

		tmp := data[i+1]
		for j := i; j >= 0; j-- {
			if tmp < data[j] {
				data[j+1] = data[j]
				data[j] = tmp
			}
		}
	}
}
